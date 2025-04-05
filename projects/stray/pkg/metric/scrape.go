package metric

import (
	"time"

	kitlog "github.com/go-kit/log"
	"github.com/go-kit/log/level"
	promCommonCfg "github.com/prometheus/common/config"
	"github.com/prometheus/common/model"
	"github.com/prometheus/prometheus/config"
	"github.com/prometheus/prometheus/scrape"
	"github.com/samber/lo"
)

type (
	PromScrape struct {
		logger        kitlog.Logger
		promDiscovery *PromDiscovery
		scrapeMgr     *scrape.Manager
	}

	PromScrapeOptions struct {
		Logger    kitlog.Logger
		Storage   *PromStorage
		Discovery *PromDiscovery
	}

	ScrapeJob struct {
		JobName     string
		Scheme      string
		MetricsPath string
		Interval    time.Duration
		Timeout     time.Duration
	}
)

func NewPromScrape(opts PromScrapeOptions) (*PromScrape, error) {
	mgr := scrape.NewManager(&scrape.Options{},
		kitlog.With(opts.Logger, LOG_COMPONENT_KEY, COMPONENT_SCRAPE),
		opts.Storage)
	return &PromScrape{
		logger:        opts.Logger,
		scrapeMgr:     mgr,
		promDiscovery: opts.Discovery}, nil
}

func (s *PromScrape) Get() *scrape.Manager {
	return s.scrapeMgr
}

func (s *PromScrape) Run() error {
	_ = level.Debug(s.logger).Log("msg", "launching scrape component")
	discoveryMgr := s.promDiscovery.Get()
	return s.scrapeMgr.Run(discoveryMgr.SyncCh())
}

func (s *PromScrape) ApplyJobs(jobs []ScrapeJob) error {
	scrapCfg, err := config.Load("", false, s.logger)
	if err != nil {
		_ = level.Error(s.logger).Log("msg", "make scrape jobs config", "err", err)
		return err
	}

	_ = level.Debug(s.logger).Log("msg", "make scrape jobs config", "jobs", jobs)
	scrapeJobs := make([]*config.ScrapeConfig, 0, len(jobs))
	lo.ForEach(jobs, func(job ScrapeJob, _ int) {
		scrapeJobs = append(scrapeJobs, &config.ScrapeConfig{
			Scheme:         job.Scheme,
			MetricsPath:    job.MetricsPath,
			JobName:        job.JobName,
			ScrapeInterval: model.Duration(job.Interval),
			ScrapeTimeout:  model.Duration(job.Timeout),
			HTTPClientConfig: promCommonCfg.HTTPClientConfig{
				FollowRedirects: true,
				EnableHTTP2:     true,
				TLSConfig:       promCommonCfg.TLSConfig{InsecureSkipVerify: true},
			},
		})
	})
	_ = level.Debug(s.logger).Log("msg", "scrape job is updating", "jobs", scrapeJobs)
	scrapCfg.ScrapeConfigs = scrapeJobs
	err = s.scrapeMgr.ApplyConfig(scrapCfg)
	if err != nil {
		_ = level.Error(s.logger).Log("msg", "scrape apply config", "err", err)
	}
	return err
}
