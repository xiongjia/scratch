package metric

import (
	"time"

	kitlog "github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/prometheus/model/labels"
	"github.com/prometheus/prometheus/storage"
	"github.com/prometheus/prometheus/tsdb"
	"github.com/prometheus/prometheus/tsdb/chunkenc"
)

type (
	promSeriesSetAdapter struct {
		log       kitlog.Logger
		seriesSet storage.SeriesSet
	}

	storageQuerierAdapter struct {
		log kitlog.Logger
		q   storage.Querier
	}
)

func makeStorageQuerierAdapter(log kitlog.Logger, q storage.Querier) storage.Querier {
	return &storageQuerierAdapter{q: q, log: log}
}

func (sq *storageQuerierAdapter) Select(sortSeries bool, hints *storage.SelectHints, matchers ...*labels.Matcher) storage.SeriesSet {
	_ = level.Debug(sq.log).Log("msg", "select", "sortSeries", sortSeries, "hints", hints, "m", matchers)
	if hints != nil {
		_ = level.Debug(sq.log).Log("msg", "hints",
			"begin", time.UnixMilli(hints.Start).Format("2006-01-02 15:04:05"),
			"end", time.UnixMilli(hints.End).Format("2006-01-02 15:04:05"),
			"step", hints.Step, "func", hints.Func, "range", hints.Range,
			"Grouping", hints.Grouping, "by", hints.By,
			"DisableTrimming", hints.DisableTrimming)
	}
	return makePromSeriesSetAdapter(sq.q.Select(sortSeries, hints, matchers...), sq.log)
}

func (sq *storageQuerierAdapter) LabelValues(name string, matchers ...*labels.Matcher) ([]string, storage.Warnings, error) {
	_ = level.Debug(sq.log).Log("msg", "labVal", "name", name, "m", matchers)
	return sq.q.LabelValues(name, matchers...)
}

func (sq *storageQuerierAdapter) LabelNames(matchers ...*labels.Matcher) ([]string, storage.Warnings, error) {
	_ = level.Debug(sq.log).Log("msg", "labNames", "m", matchers)
	return sq.q.LabelNames(matchers...)
}

func (sq *storageQuerierAdapter) Close() error {
	_ = level.Debug(sq.log).Log("msg", "close")
	return sq.q.Close()
}

func makePromSeriesSetAdapter(src storage.SeriesSet, log kitlog.Logger) storage.SeriesSet {
	return &promSeriesSetAdapter{seriesSet: src, log: log}
}

func (s *promSeriesSetAdapter) Next() bool {
	return s.seriesSet.Next()
}

func (s *promSeriesSetAdapter) At() storage.Series {
	ret := s.seriesSet.At()

	l := ret.Labels()
	_ = level.Debug(s.log).Log("msg", "at", "l", l)

	var chkIter chunkenc.Iterator
	chkIter = ret.Iterator(chkIter)
	for chkIter.Next() == chunkenc.ValFloat {
		ts, v := chkIter.At()
		_ = level.Debug(s.log).Log("msg", "at", "ts", ts, "v", v)
	}
	return ret
}

func (s *promSeriesSetAdapter) Err() error {
	return s.seriesSet.Err()
}

func (s *promSeriesSetAdapter) Warnings() storage.Warnings {
	return s.seriesSet.Warnings()
}

func createFsStorage(dbPath string, dbStats *tsdb.DBStats, l kitlog.Logger) (storage.Storage, error) {
	db, err := tsdb.Open(dbPath, l, prometheus.DefaultRegisterer, tsdb.DefaultOptions(), dbStats)
	if err != nil {
		return nil, err
	}
	return db, nil
}
