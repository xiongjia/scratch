package squirtle

import (
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

func NewMetric() *Metric {
	reg := prometheus.NewRegistry()
	handler := promhttp.HandlerFor(reg, promhttp.HandlerOpts{
		ErrorHandling: promhttp.ContinueOnError,
	})
	m := &Metric{Reg: reg, HttpHandler: handler}
	slog.Debug("created a new prometheus registry")
	return m
}

func (m *Metric) errorLogger(v ...interface{}) {

}

func (m *Metric) Factory() promauto.Factory {
	return promauto.With(m.Reg)
}
