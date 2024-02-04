package bulbasaur

import (
	"fmt"
	"log/slog"
	"stray/pkg/dugtrio"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/prometheus/scrape"
	"github.com/prometheus/prometheus/storage"
)

type EngineOptions struct {
	Logger *slog.Logger
}

type Engine struct {
	logger    *slog.Logger
	storage   *LocalStorageEngine
	scrapeMgr *scrape.Manager
}

func (e *Engine) Test() {
	fmt.Printf("test\n")
	pools := e.scrapeMgr.ScrapePools()
	fmt.Printf("Pools = %v\n", pools)

	e.scrapeMgr.Run()
}

func NewEngine(opts EngineOptions) (*Engine, error) {
	// Local storage engine
	storageOpt := &LocalStorageOptions{
		Dir:    "c:/wrk/tmp/tsdb1",
		Logger: opts.Logger,
	}
	storageEng, err := MakeLocalStorageEngine(storageOpt)
	if err != nil {
		return nil, err
	}

	fanoutStorage := storage.NewFanout(nil, storageEng.DB)
	scrapeMgr, err := scrape.NewManager(nil, dugtrio.NewKitLoggerAdapterSlog(opts.Logger), fanoutStorage, prometheus.DefaultRegisterer)
	if err != nil {
		// TODO close storages
		return nil, err
	}

	return &Engine{
		logger:    opts.Logger,
		storage:   storageEng,
		scrapeMgr: scrapeMgr,
	}, nil
}
