package metric

import (
	"log/slog"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/common/model"
	"github.com/prometheus/prometheus/config"
	"github.com/prometheus/prometheus/scrape"
	"github.com/prometheus/prometheus/storage"
	"github.com/prometheus/prometheus/util/logging"
	"github.com/samber/lo"
)

type (
	PromScrape struct {
		compLog       *slog.Logger
		promDiscovery *PromDiscovery
		scrapeMgr     *scrape.Manager
	}

	PromScrapeOptions struct {
		Log       *slog.Logger
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

func makeScrapeFanout(opts *PromScrapeOptions) storage.Storage {
	return storage.NewFanout(makeComponentLog(opts.Log, COMPONENT_STORAGE), opts.Storage)
}

func NewPromScrape(opts PromScrapeOptions) (*PromScrape, error) {
	scrapeCompLog := makeComponentLog(opts.Log, COMPONENT_SCRAPE)
	mgr, err := scrape.NewManager(
		&scrape.Options{},
		scrapeCompLog,
		logging.NewJSONFileLogger,
		makeScrapeFanout(&opts),
		prometheus.DefaultRegisterer)
	if err != nil {
		slog.Error("create scrape manager", slog.Any("err", err))
		return nil, err
	}

	return &PromScrape{
		scrapeMgr:     mgr,
		promDiscovery: opts.Discovery,
		compLog:       scrapeCompLog}, nil
}

func (s *PromScrape) Run() error {
	slog.Debug("promScrape started")
	discoveryMgr := s.promDiscovery.Get()
	return s.scrapeMgr.Run(discoveryMgr.SyncCh())
}

func (s *PromScrape) ApplyJobs(jobs []ScrapeJob) error {
	scrapCfg, err := config.Load("", s.compLog)
	if err != nil {
		slog.Error("make scrape jobs config", slog.Any("err", err))
		return err
	}

	scrapeJobs := make([]*config.ScrapeConfig, 0, len(jobs))
	lo.ForEach(jobs, func(job ScrapeJob, _ int) {
		scrapeJobs = append(scrapeJobs, &config.ScrapeConfig{
			Scheme:         job.Scheme,
			MetricsPath:    job.MetricsPath,
			JobName:        job.JobName,
			ScrapeInterval: model.Duration(job.Interval),
			ScrapeTimeout:  model.Duration(job.Timeout),
		})
	})
	slog.Debug("scrape job is updating", slog.Any("jobs", scrapeJobs))
	scrapCfg.ScrapeConfigs = scrapeJobs
	err = s.scrapeMgr.ApplyConfig(scrapCfg)
	if err != nil {
		slog.Error("scrape apply config", slog.Any("err", err))
	}
	return err
}
