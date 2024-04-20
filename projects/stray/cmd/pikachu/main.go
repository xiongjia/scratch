package main

import (
	"log/slog"
	"net/http"
	"stray/pkg/dugtrio"
	"stray/pkg/pokeball"
	"stray/pkg/squirtle"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	metric = squirtle.NewMetric()
)

func main() {
	logger := dugtrio.NewSLog(dugtrio.SLogOptions{
		SLogBaseOptions: dugtrio.SLogBaseOptions{
			Level:     slog.LevelDebug,
			AddSource: true,
		},
	})
	dugtrio.SetDefaultLogger(logger)

	engine := pokeball.New(pokeball.EngineOptions{Logger: dugtrio.DefaultLogger()})
	engine.Get("/metrics", metric.HttpHandler.ServeHTTP)

	logger.Info("Server started on :2112")
	monitor()
	http.ListenAndServe(":2112", engine)
}

func monitor() {
	// Counter
	testCountVec := metric.Factory().NewCounterVec(prometheus.CounterOpts{
		Name: "pikachu_test_opt_total",
		Help: "The total number of processed events"},
		[]string{"sn", "vm"})
	testCount := testCountVec.WithLabelValues("disk-sn-mock-1", "uuid")

	// Gauge
	testGaugeVec := metric.Factory().NewGaugeVec(prometheus.GaugeOpts{
		Name:      "test_vol_size",
		Help:      "The total volume size",
		Namespace: "stray",
		Subsystem: "pikachu"},
		[]string{"vol", "vm"})
	testGauge := testGaugeVec.WithLabelValues("/dev/pdx1", "uuid")

	go func() {
		cnt := 0
		for {
			cnt += 2
			testCount.Inc()
			testGauge.Set(float64(cnt))
			time.Sleep(2 * time.Second)
		}
	}()
}
