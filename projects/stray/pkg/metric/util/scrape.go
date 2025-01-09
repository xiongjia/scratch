package util

import (
	"log/slog"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/common/model"
	"github.com/prometheus/prometheus/config"
	"github.com/prometheus/prometheus/scrape"
	"github.com/prometheus/prometheus/storage"
	"github.com/prometheus/prometheus/util/logging"
)

type (
	PromScrape struct {
		promLog       *slog.Logger
		promDiscovery *PromDiscovery
		mgr           *scrape.Manager
	}

	PromScrapOptions struct {
		Log       *slog.Logger
		Storage   *PromStorage
		Discovery *PromDiscovery
	}
)

func (s *PromScrape) GetMgr() *scrape.Manager {
	return s.mgr
}

func NewPromScrape(opts PromScrapOptions) (*PromScrape, error) {
	fanout := storage.NewFanout(opts.Log, opts.Storage)
	mgr, err := scrape.NewManager(
		&scrape.Options{},
		opts.Log,
		logging.NewJSONFileLogger,
		fanout,
		prometheus.DefaultRegisterer)
	if err != nil {
		// TODO Adding Log
		return nil, err
	}

	scrapCfg, err := config.Load("", opts.Log)
	if err != nil {
		// TODO adding log
		return nil, err
	}

	scrapCfg.ScrapeConfigs = []*config.ScrapeConfig{
		{
			Scheme:         "http",
			MetricsPath:    "/metrics",
			JobName:        "jobMain",
			ScrapeInterval: model.Duration(time.Second * 60),
			ScrapeTimeout:  model.Duration(time.Second * 20),
		},
	}
	mgr.ApplyConfig(scrapCfg)

	return &PromScrape{
		mgr:           mgr,
		promDiscovery: opts.Discovery,
		promLog:       opts.Log}, nil
}

func (s *PromScrape) Run() error {
	discoveryMgr := s.promDiscovery.Get()
	return s.mgr.Run(discoveryMgr.SyncCh())
}
