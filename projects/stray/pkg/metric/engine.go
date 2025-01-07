package metric

import (
	"context"
	"log/slog"
	"math"

	"stray/pkg/metric/util"

	"github.com/oklog/run"
	_ "github.com/oklog/run"
	"github.com/prometheus/prometheus/model/labels"
	"github.com/prometheus/prometheus/tsdb/chunkenc"
)

type (
	Engine struct {
		promLog       *slog.Logger
		promDiscovery *util.PromDiscovery
		promScrape    *util.PromScrape
		promStorage   *util.PromStorage
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
	}, nil
}

func (eng *Engine) QueryTest() {
	querier, err := eng.promStorage.Querier(math.MinInt64, math.MaxInt64)
	if err != nil {
		return
	}

	ss := querier.Select(context.Background(), false, nil, labels.MustNewMatcher(labels.MatchEqual, "__name__", "node_cpu_seconds_total"))
	for ss.Next() {
		series := ss.At()
		slog.Debug("series", slog.String("lab", series.Labels().String()))
		it := series.Iterator(nil)
		for it.Next() == chunkenc.ValFloat {
			_, v := it.At() // We ignore the timestamp here, only to have a predictable output we can test against (below)
			slog.Debug("sample", slog.Float64("val", v))
		}
		slog.Debug("itr.Err()", slog.Any("err", it.Err()))
	}
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
