package main

import (
	"log/slog"
	"stray/internal/util"
	"stray/pkg/pokeball"

	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	slog.SetDefault(util.MakeSLog(&util.SLogOptions{Level: slog.LevelDebug, AddSource: true}))
	slog.Debug("debug test")

	engine := pokeball.New()
	engine.Get("/metrics", promhttp.Handler().ServeHTTP)

	slog.Info("Server started on :2112")
	http.ListenAndServe(":2112", engine)
}
