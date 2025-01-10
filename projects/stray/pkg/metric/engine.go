package metric

import (
	"context"
	"log/slog"
	"net/url"
	"time"

	"github.com/oklog/run"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/common/model"

	"github.com/prometheus/prometheus/promql/parser"
	prom_web "github.com/prometheus/prometheus/web"
	prom_apiv1 "github.com/prometheus/prometheus/web/api/v1"
)

type (
	Engine struct {
		ctx           context.Context
		promLog       *slog.Logger
		promDiscovery *PromDiscovery
		promScrape    *PromScrape
		promStorage   *PromStorage
		promQuerier   *PromQuerier
	}

	EngineOpts struct {
		StorageFolder string

		QuerierMaxMaxSamples int
		QuerierTimeout       time.Duration
	}

	PrometheusVersion = prom_apiv1.PrometheusVersion
)

func init() {
	model.NameValidationScheme = model.UTF8Validation
}

func NewEngine(opts EngineOpts) (*Engine, error) {
	ctx := context.Background()

	// Service Discovery
	promDiscovery, err := NewPromDiscovery(ctx, PromDiscoveryOpts{})
	if err != nil {
		slog.Error("create sd", slog.Any("err", err))
		return nil, err
	}

	// PMQL
	promQL, err := NewQuerier(PromQuerierOpts{
		MaxMaxSamples: opts.QuerierMaxMaxSamples,
		Timeout:       opts.QuerierTimeout,
	})
	if err != nil {
		slog.Error("new query error", slog.Any("err", err))
		return nil, err
	}

	// TS Database storage
	promStorage, err := NewPromStorage(opts.StorageFolder)
	if err != nil {
		slog.Error("New tsdb err", slog.Any("err", err))
		return nil, err
	}

	promScrape, err := NewPromScrape(PromScrapOptions{
		Log:       slog.Default(),
		Storage:   promStorage,
		Discovery: promDiscovery,
	})
	if err != nil {
		// TODO Close the TSDB
		slog.Error("New Scrap error", slog.Any("err", err))
		return nil, err
	}

	cors, _ := compileCORSRegexString(".*")

	// API Querier
	webOpts := &prom_web.Options{
		ListenAddresses: []string{"0.0.0.0:9096"},
		ReadTimeout:     30 * time.Second,
		MaxConnections:  512,
		Context:         context.TODO(),
		Storage:         promStorage,
		LocalStorage:    promStorage,
		// TSDBDir:        dbDir,
		QueryEngine:    promQL.QueryEngine,
		ScrapeManager:  promScrape.GetMgr(),
		RuleManager:    nil,
		Notifier:       nil,
		RoutePrefix:    "/",
		EnableAdminAPI: true,
		ExternalURL: &url.URL{
			Scheme: "http",
			Host:   "localhost:9097",
			Path:   "/",
		},
		Version:    &PrometheusVersion{},
		Gatherer:   prometheus.DefaultGatherer,
		Flags:      map[string]string{},
		CORSOrigin: cors,
	}
	slog.Debug("webOpts", webOpts)
	webHandler := prom_web.New(slog.Default(), webOpts)
	webHandler.SetReady(prom_web.Ready)
	l, err := webHandler.Listeners()
	if err != nil {
		return nil, err
	}
	go func() {
		err := webHandler.Run(context.TODO(), l, "")
		if err != nil {
			panic("web handler error")
		}
	}()

	return &Engine{
		promDiscovery: promDiscovery,
		promScrape:    promScrape,
		promStorage:   promStorage,
		promQuerier:   promQL,
	}, nil
}

func (eng *Engine) ApplyDiscoveryConfig(groups ...StaticDiscoveryConfig) {
	eng.promDiscovery.ApplyStaticTargetGroup(groups...)
}

func (eng *Engine) QueryTest() {
	if eng.promQuerier == nil {
		return
	}

	pmql := eng.promQuerier.QueryEngine
	ctx := context.TODO()

	startTime, _ := time.Parse("2006-01-02 15:04:05", "2021-01-06 15:04:05")
	// endTime, _ := time.Parse("2006-01-02 15:04:05", "2025-01-10 15:04:05")

	// exp := "promhttp_metric_handler_requests_total"
	exp := "process_cpu_seconds_total{job=~\".*\"}[5000m]"
	qry, err := pmql.NewInstantQuery(ctx, eng.promStorage, nil, exp, startTime)
	if err != nil {
		return
	}
	defer qry.Close()

	result := qry.Exec(ctx)
	if result.Err != nil {
		return
	}

	val1 := result.Value
	slog.Debug("v1", slog.Any("v1", val1))

	val := result.Value.String()
	vtype := result.Value.Type()
	slog.Debug("Result", slog.String("v", val), slog.Any("vt", vtype))

	if vtype == parser.ValueTypeMatrix {
		m, err := result.Matrix()
		if err != nil {
			return
		}

		slog.Debug("Matrix", slog.Any("m", m))
	}

	v, err := result.Vector()
	if err != nil {
		return
	}
	slog.Debug("vect", slog.Any("v", v))

	// querier, err := eng.promStorage.Querier(math.MinInt64, math.MaxInt64)
	// if err != nil {
	// 	return
	// }

	// ss := querier.Select(context.Background(), false, nil, labels.MustNewMatcher(labels.MatchEqual, "__name__", "node_cpu_seconds_total"))
	// for ss.Next() {
	// 	series := ss.At()
	// 	slog.Debug("series", slog.String("lab", series.Labels().String()))
	// 	it := series.Iterator(nil)
	// 	for it.Next() == chunkenc.ValFloat {
	// 		_, v := it.At() // We ignore the timestamp here, only to have a predictable output we can test against (below)
	// 		slog.Debug("sample", slog.Float64("val", v))
	// 	}
	// 	slog.Debug("itr.Err()", slog.Any("err", it.Err()))
	// }
}

func (eng *Engine) GetServiceDiscovery() *PromDiscovery {
	return eng.promDiscovery
}

func (eng *Engine) GetScrape() *PromScrape {
	return eng.promScrape
}

func (eng *Engine) Run() error {
	slog.Debug("Engine running")

	var runGroup run.Group

	// Service Discovery
	runGroup.Add(func() error {
		slog.Debug("discovery component started")
		return eng.promDiscovery.Run()
	}, func(err error) {
		slog.Debug("discovery component exited")
	})

	// Scrape
	runGroup.Add(func() error {
		slog.Debug("scrape component started")
		return eng.promScrape.Run()
	}, func(err error) {
		slog.Debug("scrape component exited")
	})

	return runGroup.Run()
}
