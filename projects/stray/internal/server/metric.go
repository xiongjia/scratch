package server

import (
	"net/http"
	"stray/pkg/metric"
)

type (
	EngineMetric struct {
		metric *metric.Metric
	}
)

func NewEngineMetric(mux *http.ServeMux) *EngineMetric {
	metric := metric.NewMetric()
	metric.Bind(mux, PREFIX_METRIC)
	return &EngineMetric{metric: metric}
}
