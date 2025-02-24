package main

import (
	"log/slog"
	"net/http"
	"time"

	"metric/pkg/prom"
)

func initCollector(mux *http.ServeMux) error {
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

func main() {
	serverAddress := ":3001"

	slog.SetLogLoggerLevel(slog.LevelDebug)
	slog.Debug("Test Server", slog.String("addr", serverAddress))

	// Make test server
	mux := http.NewServeMux()

	err := initCollector(mux)
	if err != nil {
		slog.Error("new collector", slog.Any("err", err))
		return
	}

	wg := prom.NewWaitGroup()
	wg.Go(func() {
		err := http.ListenAndServe(serverAddress, mux)
		if err != nil {
			slog.Error("http server error", slog.Any("err", err))
		}
	})

	<-time.After(10 * time.Second)

	wg.Wait()
}
