package metric

import (
	"log/slog"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/prometheus/tsdb"
)

func OpenLocalTsdb(dir string, reg prometheus.Registerer, opts *tsdb.Options, stats *tsdb.DBStats) (*tsdb.DB, error) {
	return tsdb.Open(dir, slog.Default().With("component", "tsdb"), reg, opts, stats)
}
