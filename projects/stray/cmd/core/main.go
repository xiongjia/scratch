package main

import (
	"log/slog"
	"os"
	"os/signal"
	"time"

	"stray/internal/server"
	"stray/pkg/metric"
	"stray/pkg/util"
)

type (
	Proc struct {
		Engine *metric.Engine
	}
)

func (p *Proc) ProcessShutdown() {
	if p.Engine != nil {
		err := p.Engine.Shutdown()
		if err != nil {
			slog.Error("metric engine shutdown", slog.Any("err", err))
		}
	}
}

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
		if p != nil {
			p.ProcessShutdown()
		}
		<-time.After(5 * time.Second)
		os.Exit(0)
	}()
}

func main() {
	var proc Proc

	procInit(&proc)

	// REST API Server
	serv, serveMux, err := server.NewServer(server.ServerConfig{
		EnableApiDoc: true,
	})
	if err != nil {
		slog.Error("Server init error", slog.Any("err", err))
		return
	}

	// prometheus engine
	eng, err := metric.NewEngine(metric.EngineOpts{
		StorageFolder:        "C:/wrk/tmp/tsdb2",
		Disable:              false,
		HttpHandler:          serv,
		ServeMux:             serveMux,
		QuerierMaxMaxSamples: 9999999999999,
	})
	if err != nil {
		slog.Error("new engine", slog.Any("err", err))
		return
	}
	proc.Engine = eng

	wg := util.NewWaitGroup()
	wg.Go(func() {
		eng.Run()
	})
	wg.Go(func() {
		slog.Info("API Server Started (0.0.0.0:8897)")
		err = server.StartServer("0.0.0.0", 8897, serv)
		if err != nil {
			slog.Error("Server error", slog.Any("err", err))
		}
	})

	// Update service discovery
	<-time.After(10 * time.Second)

	// TODO load it from cluster

	// update jobs
	err = eng.ApplyScrapeJobs([]metric.ScrapeJob{
		{
			JobName:     "jobNode1",
			Scheme:      "http",
			MetricsPath: "/metrics",
			Interval:    time.Second * 60,
			Timeout:     time.Second * 20,
		},
	})
	if err != nil {
		slog.Error("scrape job apply config", slog.Any("err", err))
	}

	// update target
	err = eng.ApplyDiscoveryConfig([]metric.StaticDiscoveryConfig{
		{
			JobName: "jobNode1",
			Targets: []metric.StaticTargetGroup{
				{
					Source:    "test",
					Addresses: []string{"172.24.6.50:9100"},
				},
			},
		},
	})
	if err != nil {
		slog.Error("service discovery apply config", slog.Any("err", err))
	}

	wg.Wait()
}
