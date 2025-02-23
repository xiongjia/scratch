package main

import (
	"log/slog"
	"net/http"

	"metric/pkg/prom"
)

func main() {
	serverAddress := ":3001"

	slog.SetLogLoggerLevel(slog.LevelDebug)
	slog.Debug("Test Server", slog.String("addr", serverAddress))

	// Make test server
	mux := http.NewServeMux()

	// Collector
	collector, err := prom.NewCollector(prom.PromCollectorConfig{Logger: prom.NewSLogAdapterHandler()})
	if err != nil {
		slog.Error("new collector", slog.Any("err", err))
		return
	}
	mux.Handle("/metric/", collector.CollectorHandler())

	wg := prom.NewWaitGroup()
	wg.Go(func() {
		err := http.ListenAndServe(serverAddress, mux)
		if err != nil {
			slog.Error("http server error", slog.Any("err", err))
		}
	})
	wg.Wait()
}
