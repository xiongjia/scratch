package metric

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type (
	Metric struct {
		reg         *prometheus.Registry
		HttpHandler http.Handler
	}

	promLoggerHandler struct{}
)

func (promLoggerHandler) Println(v ...interface{}) {
	msg := fmt.Sprint(v...)
	if len(msg) <= 0 {
		return
	}
	slog.Error("prometheus http handler error", slog.String("err", msg))
}

func NewMetric() *Metric {
	reg := prometheus.NewRegistry()
	handler := promhttp.HandlerFor(reg, promhttp.HandlerOpts{
		ErrorHandling: promhttp.ContinueOnError,
		ErrorLog:      &promLoggerHandler{},
	})
	return &Metric{reg: reg, HttpHandler: handler}
}

func (m *Metric) Bind(mux *http.ServeMux, router string) {

}

func (m *Metric) Factory() promauto.Factory {
	return promauto.With(m.reg)
}
