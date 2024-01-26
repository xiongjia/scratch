package main

import (
	"log/slog"
	"net/http"
	"stray/internal/util"
	"stray/pkg/pokeball"
	"stray/pkg/squirtle"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	metric = squirtle.NewMetric()
)

func main() {
	slog.SetDefault(util.MakeSLog(&util.SLogOptions{Level: slog.LevelDebug, AddSource: true}))
	slog.Debug("debug test")

	engine := pokeball.New()
	engine.Get("/metrics", metric.HttpHandler.ServeHTTP)

	slog.Info("Server started on :2112")
	monitor()
	http.ListenAndServe(":2112", engine)
}

func monitor() {
	// TODO testing other counters
	testCount := metric.Factory().NewCounter(prometheus.CounterOpts{
		Name: "pikachu_test_opt_total",
		Help: "The total number of processed events",
	})
	go func() {
		for {
			testCount.Inc()
			time.Sleep(2 * time.Second)
		}
	}()
}
