package squirtle

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Metric struct {
	Reg         *prometheus.Registry
	HttpHandler http.Handler
}

type promhttpErrorLogger struct{ promhttp.Logger }

func (l *promhttpErrorLogger) Println(v ...interface{}) {
	slog.Error("prometheus client error", slog.String("client-error", fmt.Sprint(v...)))
}

func NewMetric() *Metric {
	reg := prometheus.NewRegistry()
	handler := promhttp.HandlerFor(reg, promhttp.HandlerOpts{
		ErrorHandling: promhttp.ContinueOnError,
		ErrorLog:      &promhttpErrorLogger{},
	})
	m := &Metric{Reg: reg, HttpHandler: handler}
	slog.Debug("created a new prometheus registry")
	return m
}

func (m *Metric) Factory() promauto.Factory {
	return promauto.With(m.Reg)
}
