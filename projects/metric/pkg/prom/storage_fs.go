package prom

import (
	kitlog "github.com/go-kit/log"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/prometheus/storage"
	"github.com/prometheus/prometheus/tsdb"
)

func createFsStorage(dbPath string, dbStats *tsdb.DBStats, l kitlog.Logger) (storage.Storage, error) {
	db, err := tsdb.Open(dbPath, l, prometheus.DefaultRegisterer, tsdb.DefaultOptions(), dbStats)
	if err != nil {
		return nil, err
	}
	return db, nil
}
