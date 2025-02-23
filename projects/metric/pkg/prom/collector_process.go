package prom

import (
	kitlog "github.com/go-kit/log"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
)

type (
	ProcessCollector struct {
		logger        kitlog.Logger
		procCollector prometheus.Collector
	}
)

func init() {
	registerCollectorInner("Rhino.Process", newProcessCollector)
}

func (c *ProcessCollector) Update(ch chan<- prometheus.Metric) error {
	c.procCollector.Collect(ch)
	return nil
}

func newProcessCollector(logger kitlog.Logger) (Collector, error) {
	return &ProcessCollector{
		logger:        logger,
		procCollector: collectors.NewProcessCollector(collectors.ProcessCollectorOpts{Namespace: namespace}),
	}, nil
}
