package metric

import (
	"context"
	"log/slog"

	"stray/pkg/metric/util"

	"github.com/oklog/run"
	_ "github.com/oklog/run"
	"github.com/prometheus/prometheus/tsdb"
)

type (
	Engine struct {
		promLog       *slog.Logger
		promDiscovery *util.PromDiscovery
		promScrape    *util.PromScrape
		promStorage   *util.PromStorage
	}
)

func NewEngine() (*Engine, error) {
	promLog := slog.Default()
	promDiscovery, err := util.NewPromDiscovery(context.TODO())
	if err != nil {
		slog.Error("New service discovery", slog.Any("err", err))
		return nil, err
	}

	promStorage := util.NewPromStorage(tsdb.NewDBStats())
	promScrape, err := util.NewPromScrape(util.PromScrapOptions{
		Log:       promLog.With("component", "scrape"),
		Storage:   promStorage,
		Discovery: promDiscovery,
	})
	if err != nil {
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
