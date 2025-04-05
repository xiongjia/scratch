package metric

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync"
	"time"

	"crypto/tls"

	kitlog "github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/oklog/run"
)

type (
	EngineOptions struct {
		Logger LogAdapterHandler

		StorageDBPath string

		QuerierMaxMaxSamples int
		QuerierTimeout       time.Duration

		MasterNode        bool
		MasterNodeApiAddr string
		MasterNodeApiPort uint16
	}

	Engine struct {
		mtx    sync.Mutex
		logger kitlog.Logger

		ctx     context.Context
		disable bool

		promDiscovery *PromDiscovery
		promStorage   *PromStorage
		promQuerier   *PromQuerier
		promScrape    *PromScrape
		promEngApi    *EngineAPI

		targetGroups []StaticDiscoveryConfig
		scrapeJobs   []ScrapeJob

		currentMasterNode bool
		masterNodeApiAddr string
		masterNodeApiPort uint16
	}

	PromTargetsJobs struct {
		Jobs    []ScrapeJob
		Targets []StaticDiscoveryConfig
	}
)

func NewEngine(opts EngineOptions) (*Engine, error) {
	promLog := NewLogAdapter(opts.Logger)
	engCtx := context.Background()

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
		FsTsdbPath: opts.StorageDBPath,
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
		_ = level.Error(promLog).Log("msg", "create api", "err", err)
		return nil, err
	}

	eng := &Engine{
		logger:        promLog,
		ctx:           engCtx,
		promDiscovery: promDiscovery,
		promQuerier:   promQuerier,
		promScrape:    promScrape,
		promEngApi:    promEngApi,
		targetGroups:  []StaticDiscoveryConfig{},
		scrapeJobs:    []ScrapeJob{},
		promStorage:   promStorage,

		currentMasterNode: opts.MasterNode,
	}
	return eng, nil
}

func (e *Engine) GetQuerier() *PromQuerier {
	return e.promQuerier
}

func (e *Engine) GetStorage() *PromStorage {
	e.mtx.Lock()
	defer e.mtx.Unlock()
	return e.promStorage
}

func (e *Engine) ApplyDiscoveryConfig(groups []StaticDiscoveryConfig) error {
	if e.disable {
		_ = level.Info(e.logger).Log("msg", "ignored service discovery config (engine is disabled)", "group", groups)
		return nil
	}
	e.mtx.Lock()
	defer e.mtx.Unlock()
	if compareStaticTargetConf(groups, e.targetGroups) {
		// No change
		return nil
	}
	err := e.promDiscovery.ApplyStaticTargetGroup(groups)
	if err != nil {
		_ = level.Error(e.logger).Log("msg", "apply static group", "err", err, "grp", groups)
		return err
	}
	e.targetGroups = groups
	return nil
}

func (e *Engine) ApplyScrapeJobs(jobs []ScrapeJob) error {
	if e.disable {
		_ = level.Info(e.logger).Log("msg", "ignored scrape jobs (engine is disabled)", "jobs", jobs)
		return nil
	}
	e.mtx.Lock()
	defer e.mtx.Unlock()
	if compareScrapeJobs(jobs, e.scrapeJobs) {
		// No change
		return nil
	}
	err := e.promScrape.ApplyJobs(jobs)
	if err != nil {
		_ = level.Error(e.logger).Log("msg", "apply scrape jobs", "err", err, "jobs", jobs)
		return err
	}
	e.scrapeJobs = jobs
	return nil
}

func (eng *Engine) Register(prefix string) (http.Handler, error) {
	return eng.promEngApi.Register(prefix)
}

func (e *Engine) Run() error {
	_ = level.Debug(e.logger).Log("msg", "starting prometheus")
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

func (e *Engine) ClearTargetsJobs() {
	err := e.ApplyTargetsJobs(PromTargetsJobs{Jobs: []ScrapeJob{}, Targets: []StaticDiscoveryConfig{}})
	if err != nil {
		_ = level.Error(e.logger).Log("msg", "clean target job", "err", err)
	}
}

func (e *Engine) ApplyTargetsJobs(src PromTargetsJobs) error {
	err := e.ApplyScrapeJobs(src.Jobs)
	if err != nil {
		return err
	}
	return e.ApplyDiscoveryConfig(src.Targets)
}

func (e *Engine) SetMasterApi(addr string, port uint16) {
	e.mtx.Lock()
	defer e.mtx.Unlock()
	e.masterNodeApiAddr = addr
	e.masterNodeApiPort = port
}

func (e *Engine) SetMaster(currentMaster bool) {
	e.mtx.Lock()
	defer e.mtx.Unlock()
	e.currentMasterNode = currentMaster
}

func (e *Engine) IsCurrentMaster() bool {
	e.mtx.Lock()
	defer e.mtx.Unlock()
	return e.currentMasterNode
}

func (e *Engine) makeMasterUrl() (*url.URL, error) {
	e.mtx.Lock()
	defer e.mtx.Unlock()
	if len(e.masterNodeApiAddr) == 0 {
		return nil, errors.New("master not ready")
	}
	return url.Parse(fmt.Sprintf("https://%s:%d", e.masterNodeApiAddr, e.masterNodeApiPort))
}

func (e *Engine) ForwardHandler(ctx context.Context, w http.ResponseWriter, r *http.Request) bool {
	if e.IsCurrentMaster() {
		return false
	}

	targetURL, err := e.makeMasterUrl()
	if err != nil {
		_ = level.Error(e.logger).Log("msg", "prom master not ready")
		return false
	}
	proxy := httputil.NewSingleHostReverseProxy(targetURL)
	proxy.Transport = &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}
	proxy.ServeHTTP(w, r)
	return true
}
