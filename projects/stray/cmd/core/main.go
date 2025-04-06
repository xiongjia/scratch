package main

import (
	"log/slog"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"

	"stray/pkg/metric"
	"stray/pkg/util"
)

type (
	Proc struct {
		Collector *metric.MetricCollector
		Engine    *metric.Engine
	}
)

func procInit(p *Proc) {
	util.InitDefaultLog(&util.LogOption{
		Level:     slog.LevelDebug,
		AddSource: false,
	})
	util.DebugDumpDeps()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt)
	go func() {
		sig := <-sigs
		slog.Debug("got os interrupt", slog.String("sig", sig.String()))
		// if p != nil {
		// 	p.ProcessShutdown()
		// }
		// <-time.After(5 * time.Second)
		os.Exit(0)
	}()
}

func Touch(eng *metric.Engine) {
	// update jobs
	if err := eng.ApplyScrapeJobs([]metric.ScrapeJob{
		{
			JobName:     "jobNode1",
			Scheme:      "http",
			MetricsPath: "/metric",
			Interval:    time.Second * 60,
			Timeout:     time.Second * 20,
		},
	}); err != nil {
		slog.Error("scrape job apply config", slog.Any("err", err))
	}

	// update target
	if err := eng.ApplyDiscoveryConfig([]metric.StaticDiscoveryConfig{
		{
			JobName: "jobNode1",
			Targets: []metric.StaticTargetGroup{
				{
					Source:    "test",
					Addresses: []string{"127.0.0.1:8897"},
				},
			},
		},
	}); err != nil {
		slog.Error("service discovery apply config", slog.Any("err", err))
	}

	// load rules
	eng.LoadRuleGroups()
}

func main() {
	var proc Proc
	procInit(&proc)

	collector, err := metric.NewMetricCollector(metric.MetricCollectorConfig{Logger: metric.NewSLogAdapterHandler()})
	if err != nil {
		slog.Error("new collector", slog.Any("err", err))
		return
	}
	proc.Collector = collector

	eng, err := metric.NewEngine(metric.EngineOptions{
		Logger:               metric.NewSLogAdapterHandler(),
		StorageDBPath:        "C:/wrk/tmp/tsdb2",
		QuerierMaxMaxSamples: 9999999999999,
		QuerierTimeout:       30 * time.Second,
	})
	if err != nil {
		slog.Error("new engine", slog.Any("err", err))
		return
	}
	proc.Engine = eng
	mux := http.NewServeMux()
	mux.Handle("/metric/", proc.Collector.CollectorHandler())
	engApiHandler, _ := proc.Engine.Register("/")
	mux.Handle("/mon/{p...}", http.StripPrefix("/mon", engApiHandler))

	wg := util.NewWaitGroup()
	wg.Go(func() { eng.Run() })
	wg.Go(func() {
		slog.Info("API Server Started")
		addr := net.JoinHostPort("0.0.0.0", strconv.Itoa(8897))
		httpServer := &http.Server{Handler: mux, Addr: addr}
		err := httpServer.ListenAndServe()
		if err != nil {
			slog.Error("Server error", slog.Any("err", err))
		}
	})

	<-time.After(10 * time.Second)
	Touch(eng)
	wg.Wait()
}
