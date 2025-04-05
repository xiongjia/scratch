package metric

import (
	"fmt"
	"net/http"
	"sync"

	kitlog "github.com/go-kit/log"
	"github.com/go-kit/log/level"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type (
	Collector interface {
		Update(ch chan<- prometheus.Metric) error
	}

	NodeCollector struct {
		CollectorName string
		Enabled       bool
		Metric        Collector
	}

	NodeCollectors struct {
		collectors []*NodeCollector
		logger     kitlog.Logger
	}

	MetricCollectorConfig struct {
		Logger LogAdapterHandler
	}

	MetricCollector struct {
		promRegistry *prometheus.Registry
		httpHandler  http.Handler
		logger       kitlog.Logger
	}

	httpErrorLogger struct {
		promhttp.Logger
		logger kitlog.Logger
	}
)

var (
	collectorsMtx           = sync.Mutex{}
	collectorFactoriesInner = make(map[string]func(kitlog.Logger) (Collector, error))
	collectorFactories      = make(map[string]func() (Collector, error))
)

func RegisterCollector(collector string, factory func() (Collector, error)) {
	collectorsMtx.Lock()
	defer collectorsMtx.Unlock()
	collectorFactories[collector] = factory
}

func registerCollectorInner(collector string, factory func(kitlog.Logger) (Collector, error)) {
	collectorsMtx.Lock()
	defer collectorsMtx.Unlock()
	collectorFactoriesInner[collector] = factory
}

func PushMetric(ch chan<- prometheus.Metric, desc *prometheus.Desc,
	valueType prometheus.ValueType, value float64, labelValues ...string) error {
	metric, err := prometheus.NewConstMetric(desc, valueType, value, labelValues...)
	if err != nil {
		return err
	}
	ch <- metric
	return nil
}

func (*NodeCollectors) Describe(chan<- *prometheus.Desc) {
	// NOP
}

func (l *httpErrorLogger) Println(v ...interface{}) {
	msg := fmt.Sprint(v...)
	if len(msg) <= 0 {
		return
	}
	_ = level.Error(l.logger).Log("msg", msg)
}

func NewMetricCollector(cfg MetricCollectorConfig) (*MetricCollector, error) {
	promLog := NewLogAdapter(cfg.Logger)
	nodeCollectors, err := newNodeCollectors(promLog)
	if err != nil {
		_ = level.Error(promLog).Log("msg", "creating node collector failed", "err", err)
		return nil, err
	}
	reg := prometheus.NewRegistry()
	err = reg.Register(nodeCollectors)
	if err != nil {
		_ = level.Error(promLog).Log("msg", "node collector register error", "err", err)
		return nil, err
	}
	handler := promhttp.HandlerFor(reg, promhttp.HandlerOpts{
		ErrorHandling: promhttp.ContinueOnError,
		ErrorLog:      &httpErrorLogger{logger: promLog},
	})
	return &MetricCollector{
		promRegistry: reg,
		httpHandler:  handler,
		logger:       promLog}, nil
}

func (c *MetricCollector) CollectorHandler() http.HandlerFunc {
	return c.httpHandler.ServeHTTP
}

func newNodeCollectors(logger kitlog.Logger) (*NodeCollectors, error) {
	collectorsMtx.Lock()
	defer collectorsMtx.Unlock()

	collectors := make([]*NodeCollector, 0)
	for name, factory := range collectorFactories {
		collector, err := factory()
		if err != nil {
			_ = level.Error(logger).Log("msg", "create node collector error", "err", err)
			continue
		}
		nodeCollectors := &NodeCollector{
			CollectorName: name,
			Enabled:       true,
			Metric:        collector,
		}
		collectors = append(collectors, nodeCollectors)
	}

	// inner collectors
	for name, factory := range collectorFactoriesInner {
		collector, err := factory(logger)
		if err != nil {
			_ = level.Error(logger).Log("msg", "create node collector error", "err", err)
			continue
		}
		nodeCollectors := &NodeCollector{
			CollectorName: name,
			Enabled:       true,
			Metric:        collector,
		}
		collectors = append(collectors, nodeCollectors)
	}
	return &NodeCollectors{collectors: collectors, logger: logger}, nil
}

func (c *NodeCollectors) Collect(updateChan chan<- prometheus.Metric) {
	var wg sync.WaitGroup
	for _, collector := range c.collectors {
		wg.Add(1)
		go (func(nodeCollector *NodeCollector, collectorWg *sync.WaitGroup) {
			defer collectorWg.Done()
			err := nodeCollector.Metric.Update(updateChan)
			if err != nil {
				_ = level.Error(c.logger).Log("msg", "metric updating error", "err", err)
			}
		})(collector, &wg)
	}
	wg.Wait()
}
