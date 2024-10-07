package metric_test

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"stray/pkg/metric"
	"testing"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/stretchr/testify/assert"
	"go.uber.org/goleak"
)

func TestMetric(t *testing.T) {
	defer goleak.VerifyNone(t)

	mux := http.NewServeMux()
	m := metric.NewMetric()
	m.Bind(mux, "metric")
	ts := httptest.NewServer(mux)
	defer ts.Close()

	// To register a Counter
	testCountVec := m.Factory().NewCounterVec(prometheus.CounterOpts{
		Name:      "m1_counter",
		Help:      "metric1",
		Namespace: "nm1",
	}, []string{"lab1", "lab2"})
	testCount := testCountVec.WithLabelValues("lab1-val", "lab2-val")
	testCount.Inc()

	res, err := http.Get(fmt.Sprintf("%s/metric", ts.URL))
	assert.NoError(t, err)
	defer res.Body.Close()

	content, err := io.ReadAll(res.Body)
	assert.NoError(t, err)
	assert.NotEmpty(t, content)
	t.Logf("content %s", string(content))

	// prometheus.Collector
}
