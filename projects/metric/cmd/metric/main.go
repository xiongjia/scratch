package main

import (
	"log/slog"
	"metric/pkg/prom"
	"net/http"
	"time"
)

func makeCollector(mux *http.ServeMux) error {
	collector, err := prom.NewCollector(prom.PromCollectorConfig{
		Logger: prom.NewSLogAdapterHandler(),
	})
	if err != nil {
		slog.Error("new collector", slog.Any("err", err))
		return err
	}
	mux.Handle("/metric/", collector.CollectorHandler())
	return nil
}

func touchJobs(eng *prom.Engine) {
	// jobNode1 = http://172.24.6.50:9100/metrics
	// update jobs
	err := eng.ApplyScrapeJobs([]prom.ScrapeJob{
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
		return
	}

	// Updating scrape jobs
	err = eng.ApplyDiscoveryConfig([]prom.StaticDiscoveryConfig{
		{
			JobName: "jobNode1",
			Targets: []prom.StaticTargetGroup{
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
}

func makePromEng(mux *http.ServeMux) (*prom.Engine, error) {
	engOption := prom.EngineOptions{
		Logger:  prom.NewSLogAdapterHandler(),
		Disable: false,
		// StorageType:   prom.STORAGE_FS,
		// StorageFsPath: "c:/wrk/tmp/tsdb3",
		StorageType: prom.STORAGE_DB,
		// StorageFsPath:        ":memory:",
		StorageFsPath:        "C:/wrk/tmp/tsdb-test1.db",
		QuerierMaxMaxSamples: 999999999999999,
		QuerierTimeout:       20 * time.Second,
	}

	// xxx Enable for local prom test
	engOption.StorageType = prom.STORAGE_FS
	engOption.StorageFsPath = "c:/wrk/tmp/tsdb3"

	eng, err := prom.NewEngine(engOption)
	if err != nil {
		slog.Error("new engine", slog.Any("err", err))
		return nil, err
	}
	engApiHandler, err := eng.Register("/")
	if err != nil {
		slog.Error("bind engine api", slog.Any("err", err), slog.Any("h", engApiHandler))
		return nil, err
	}
	mux.Handle("/mon/{p...}", http.StripPrefix("/mon", engApiHandler))
	return eng, nil
}

func main() {
	slog.SetLogLoggerLevel(slog.LevelDebug)

	serverAddress := ":3001"
	slog.Debug("Test Server", slog.String("addr", serverAddress))
	mux := http.NewServeMux()

	err := makeCollector(mux)
	if err != nil {
		slog.Error("new collector", slog.Any("err", err))
		return
	}

	eng, err := makePromEng(mux)
	if err != nil {
		slog.Error("new engine", slog.Any("err", err))
		return
	}

	wg := prom.NewWaitGroup()
	wg.Go(func() {
		err := http.ListenAndServe(serverAddress, mux)
		if err != nil {
			slog.Error("http server error", slog.Any("err", err))
		}
	})
	wg.Go(func() {
		err := eng.Run()
		if err != nil {
			slog.Error("engine error", slog.Any("err", err))
		}
	})

	<-time.After(10 * time.Second)
	touchJobs(eng)
	wg.Wait()
}
