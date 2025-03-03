package prom

import (
	"context"

	kitlog "github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/prometheus/prometheus/model/exemplar"
	"github.com/prometheus/prometheus/model/histogram"
	"github.com/prometheus/prometheus/model/labels"
	"github.com/prometheus/prometheus/model/metadata"
	"github.com/prometheus/prometheus/storage"
)

type (
	dbAppender struct {
		log     kitlog.Logger
		ctx     context.Context
		storage *storageDb
	}
)

func (db *dbAppender) Append(ref storage.SeriesRef, l labels.Labels, t int64, v float64) (storage.SeriesRef, error) {
	_ = level.Debug(db.log).Log("msg", "rcd appender", "l", l, "t", t, "v", v)
	return db.storage.rcdAppend(ref, l, t, v)
}

func (db *dbAppender) UpdateMetadata(ref storage.SeriesRef, l labels.Labels, m metadata.Metadata) (storage.SeriesRef, error) {
	return db.storage.rcdLabsUpdateMetadata(ref, l, &m)
}

func (db *dbAppender) Commit() error {
	// Don't support commit & rollback
	_ = level.Debug(db.log).Log("msg", "db appender commit")
	return nil
}

func (db *dbAppender) Rollback() error {
	// Don't support commit & rollback
	_ = level.Debug(db.log).Log("msg", "db appender rollback")
	return nil
}

func (dbAppender) AppendExemplar(ref storage.SeriesRef, l labels.Labels, v exemplar.Exemplar) (storage.SeriesRef, error) {
	// TODO
	return 0, nil
}

func (dbAppender) AppendHistogram(storage.SeriesRef, labels.Labels, int64, *histogram.Histogram, *histogram.FloatHistogram) (storage.SeriesRef, error) {
	// TODO
	return 0, nil
}
