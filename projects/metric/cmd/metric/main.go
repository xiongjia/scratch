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

func makePromEng(mux *http.ServeMux) (*prom.Engine, error) {
	eng, err := prom.NewEngine(prom.EngineOptions{
		Logger:               prom.NewSLogAdapterHandler(),
		Disable:              false,
		DBPath:               "c:/wrk/tmp/tsdb3",
		QuerierMaxMaxSamples: 999999999999999,
		QuerierTimeout:       20 * time.Second,
	})
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

	wg.Wait()
}
