package metric

import (
	"context"
	"log/slog"
	"time"

	"stray/pkg/metric/util"

	"github.com/oklog/run"
	"github.com/prometheus/prometheus/promql/parser"
)

type (
	Engine struct {
		promLog       *slog.Logger
		promDiscovery *util.PromDiscovery
		promScrape    *util.PromScrape
		promStorage   *util.PromStorage
		promQuerier   *util.PromQuerier
	}

	EngineOpts struct {
		StorageFolder string
	}
)

func NewEngine(opts EngineOpts) (*Engine, error) {
	promLog := slog.Default()

	promDiscovery, err := util.NewPromDiscovery(context.TODO())
	if err != nil {
		slog.Error("New service discovery", slog.Any("err", err))
		return nil, err
	}

	promStorage, err := util.NewPromStorage(opts.StorageFolder)
	if err != nil {
		slog.Error("New tsdb err", slog.Any("err", err))
		return nil, err
	}

	promQL, err := util.NewQuerier()
	if err != nil {
		slog.Error("new query error", slog.Any("err", err))
		return nil, err
	}

	promScrape, err := util.NewPromScrape(util.PromScrapOptions{
		Log:       promLog.With("component", "scrape"),
		Storage:   promStorage,
		Discovery: promDiscovery,
	})
	if err != nil {
		// TODO Close the TSDB
		slog.Error("New Scrap error", slog.Any("err", err))
		return nil, err
	}

	return &Engine{
		promLog:       promLog,
		promDiscovery: promDiscovery,
		promScrape:    promScrape,
		promStorage:   promStorage,
		promQuerier:   promQL,
	}, nil
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

func (eng *Engine) GetServiceDiscovery() *util.PromDiscovery {
	return eng.promDiscovery
}

func (eng *Engine) GetScrape() *util.PromScrape {
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
