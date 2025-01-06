package main

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"stray/internal/server"
	"stray/pkg/metric"
	metricUtil "stray/pkg/metric/util"
	"stray/pkg/util"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/common/model"
	"github.com/prometheus/prometheus/config"
	"github.com/prometheus/prometheus/discovery/targetgroup"
	"github.com/prometheus/prometheus/scrape"
	"github.com/prometheus/prometheus/storage"
	"github.com/prometheus/prometheus/tsdb"
	"github.com/prometheus/prometheus/util/logging"

	"github.com/prometheus/prometheus/discovery"
)

func staticConfig(addrs ...string) discovery.StaticConfig {
	var cfg discovery.StaticConfig
	for i, addr := range addrs {
		cfg = append(cfg, &targetgroup.Group{
			Source: fmt.Sprint(i),
			Targets: []model.LabelSet{
				{model.AddressLabel: model.LabelValue(addr)},
			},
		})
	}
	return cfg
}

func procInit() {
	util.InitDefaultLog(&util.LogOption{
		Level:     slog.LevelDebug,
		AddSource: false,
	})
	util.DebugDumpDeps()

	log := slog.Default()
	// logLvl := &promslog.AllowedLevel{}
	// logLvl.Set("debug")
	// promLog := promslog.New(&promslog.Config{Level: logLvl})
	// slog.SetDefault(promLog)
	// promLog.Debug("test")

	promLog := log.With("component", "scrape manager")

	localStorage := metricUtil.NewPromStorage(tsdb.NewDBStats())
	// remoteStorage = remote.NewStorage(logger.With("component", "remote"), prometheus.DefaultRegisterer, localStorage.StartTime, localStoragePath, time.Duration(cfg.RemoteFlushDeadline), scraper, cfg.scrape.AppendMetadata)
	fanout := storage.NewFanout(promLog, localStorage)

	scrapeMgr, err := scrape.NewManager(
		&scrape.Options{},
		promLog,
		logging.NewJSONFileLogger,
		fanout,
		prometheus.DefaultRegisterer)
	if err != nil {
		slog.Error("Scrap manger new error", slog.Any("err", err))
		return
	}

	sdMetrics, err := discovery.CreateAndRegisterSDMetrics(prometheus.DefaultRegisterer)
	if err != nil {
		slog.Error("failed to register service discovery metrics", "err", err)
		return
	}

	ctx := context.TODO()
	discoveryMgr := discovery.NewManager(ctx, promLog.With("component", "discovery"), prometheus.DefaultRegisterer, sdMetrics)
	discoveryCfg := map[string]discovery.Configs{
		"job1": {staticConfig("172.24.6.50:9100")},
	}
	discoveryMgr.ApplyConfig(discoveryCfg)
	go func() {
		discoveryMgr.Run()
	}()

	// slog.Debug("targs ", slog.Any("tags", tag))

	scrapCfg := &config.Config{}
	scrapCfg.ScrapeConfigs = []*config.ScrapeConfig{
		{
			Scheme:         "http",
			MetricsPath:    "/metrics",
			JobName:        "job1",
			ScrapeInterval: model.Duration(time.Second * 10),
			ScrapeTimeout:  model.Duration(time.Second * 3),
		},
	}
	scrapeMgr.ApplyConfig(scrapCfg)

	cfgText1 := `
scrape_configs:
 - job_name: job1
   static_configs:
   - targets: ["172.24.6.50:9100"]
`
	cfg2, err := config.Load(cfgText1, promLog)
	if err != nil {
		slog.Error("Error", slog.Any("err", err))
		return
	}
	slog.Error("Test cfg ", slog.Any("cfg2", cfg2))
	scrapeMgr.ApplyConfig(cfg2)

	// 	scrapCfgs, err := cfg2.GetScrapeConfigs()
	// 	if err != nil {
	// 		return
	// 	}

	// for _, cfgItem := range scrapCfgs {
	// 	discoveryCfg := make(map[string]discovery.Configs)
	// 	discoveryCfg["job1"] = cfgItem.ServiceDiscoveryConfigs
	// 	discoveryMgr.ApplyConfig(discoveryCfg)
	// }
	// slog.Debug("cfg", slog.Any("cfg2", cfg2), slog.Any("scrapCfg1", scrapCfgs))

	go func() {
		scrapeMgr.ScrapePools()
		scrapeMgr.Run(discoveryMgr.SyncCh())
	}()

	// targets2 := <-discoveryMgr.SyncCh()
	// fmt.Printf("Targets %v\n", targets2)
}

func main() {
	procInit()

	serv, err := server.NewServer(server.ServerConfig{
		EnableApiDoc: true,
	})
	if err != nil {
		slog.Error("Server init error", slog.Any("err", err))
		return
	}

	eng, err := metric.NewEngine()
	if err != nil {
		slog.Error("Engine init error", slog.Any("err", err))
		return
	}

	wg := util.NewWaitGroup()
	wg.Go(func() {
		slog.Info("API Server Started (0.0.0.0:8897)")
		err = server.StartServer("0.0.0.0", 8897, serv)
		if err != nil {
			slog.Error("Server error", slog.Any("err", err))
		}
	})

	wg.Go(func() {
		eng.Run()
	})

	wg.Wait()
}
