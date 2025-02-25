package prom

import (
	"context"
	"net/http"
	"time"

	kitlog "github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/oklog/run"
)

type (
	EngineOptions struct {
		Logger LogAdapterHandler

		Disable bool

		StorageType   string
		StorageFsPath string

		QuerierMaxMaxSamples int
		QuerierTimeout       time.Duration
	}

	Engine struct {
		logger kitlog.Logger

		ctx     context.Context
		disable bool

		promDiscovery *PromDiscovery
		promStorage   *PromStorage
		promQuerier   *PromQuerier
		promScrape    *PromScrape
		promEngApi    *EngineAPI
	}
)

func (e *Engine) GetQuerier() *PromQuerier {
	return e.promQuerier
}

func (e *Engine) GetStorage() *PromStorage {
	return e.promStorage
}

func (eng *Engine) Register(prefix string) (http.Handler, error) {
	return eng.promEngApi.Register(prefix)
}

func NewEngine(opts EngineOptions) (*Engine, error) {
	engCtx := context.TODO()
	promLog := NewLogAdapter(opts.Logger)

	// Service Discovery
	promDiscovery, err := NewPromDiscovery(engCtx, PromDiscoveryOpts{Logger: promLog})
	if err != nil {
		_ = level.Error(promLog).Log("msg", "create service discovery", "err", err)
		return nil, err
	}

	promQuerier, err := NewQuerier(PromQuerierOpts{
		Log:           promLog,
		MaxMaxSamples: opts.QuerierMaxMaxSamples,
		Timeout:       opts.QuerierTimeout,
	})
	if err != nil {
		_ = level.Error(promLog).Log("msg", "create pm QL engine", "err", err)
		return nil, err
	}
	promStorage, err := NewPromStorage(PromStorageOpts{
		Log:        promLog,
		Type:       opts.StorageType,
		FsTsdbPath: opts.StorageFsPath,
	})
	if err != nil {
		_ = level.Error(promLog).Log("msg", "create storage", "err", err)
		return nil, err
	}

	promScrape, err := NewPromScrape(PromScrapeOptions{
		Logger:    promLog,
		Storage:   promStorage,
		Discovery: promDiscovery,
	})
	if err != nil {
		// TODO close storage
		_ = level.Error(promLog).Log("msg", "create scrape", "err", err)
		return nil, err
	}

	promEngApi, err := NewEngineAPI(EngineAPIOpts{
		Logger:  promLog,
		Querier: promQuerier,
		Storage: promStorage,
		Scrape:  promScrape,
	})
	if err != nil {
		// TODO close storage
		_ = level.Error(promLog).Log("msg", "create api", "err", err)
		return nil, err
	}
	return &Engine{
		logger:        promLog,
		ctx:           engCtx,
		disable:       opts.Disable,
		promDiscovery: promDiscovery,
		promQuerier:   promQuerier,
		promStorage:   promStorage,
		promScrape:    promScrape,
		promEngApi:    promEngApi,
	}, nil
}

func (e *Engine) ApplyDiscoveryConfig(groups []StaticDiscoveryConfig) error {
	if e.disable {
		_ = level.Info(e.logger).Log("msg", "ignored service discovery config (engine is disabled)", "group", groups)
		return nil
	}
	return e.promDiscovery.ApplyStaticTargetGroup(groups)
}

func (e *Engine) ApplyScrapeJobs(jobs []ScrapeJob) error {
	if e.disable {
		_ = level.Info(e.logger).Log("msg", "ignored scrape jobs (engine is disabled)", "jobs", jobs)
		return nil
	}
	return e.promScrape.ApplyJobs(jobs)
}

func (e *Engine) Run() error {
	var runGroup run.Group

	// Service Discovery
	runGroup.Add(func() error {
		_ = level.Debug(e.logger).Log("msg", "discovery component started")
		return e.promDiscovery.Run()
	}, func(err error) {
		_ = level.Debug(e.logger).Log("msg", "discovery component exited")
		if err != nil {
			_ = level.Debug(e.logger).Log("msg", "discovery component  exit with err", "err", err)
		}
	})

	// Scrape
	runGroup.Add(func() error {
		_ = level.Debug(e.logger).Log("msg", "scrape component started")
		return e.promScrape.Run()
	}, func(err error) {
		_ = level.Debug(e.logger).Log("msg", "scrape component exited")
		if err != nil {
			_ = level.Error(e.logger).Log("msg", "scrape component exit with error", "err", err)
		}
	})

	return runGroup.Run()
}
