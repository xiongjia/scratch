package prom

import (
	"time"

	kitlog "github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/prometheus/prometheus/model/labels"
	"github.com/prometheus/prometheus/storage"
)

type (
	storageQuerier struct {
		log kitlog.Logger
		q   storage.Querier
	}
)

func makeStorageQuerier(log kitlog.Logger, q storage.Querier) storage.Querier {
	return &storageQuerier{q: q, log: log}
}

func (sq *storageQuerier) Select(sortSeries bool, hints *storage.SelectHints, matchers ...*labels.Matcher) storage.SeriesSet {
	_ = level.Debug(sq.log).Log("msg", "select", "sortSeries", sortSeries, "hints", hints, "m", matchers)
	if hints != nil {
		_ = level.Debug(sq.log).Log("msg", "hints",
			"begin", time.UnixMilli(hints.Start).Format("2006-01-02 15:04:05"),
			"end", time.UnixMilli(hints.End).Format("2006-01-02 15:04:05"),
			"step", hints.Step, "func", hints.Func, "range", hints.Range,
			"Grouping", hints.Grouping, "by", hints.By, "DisableTrimming", hints.DisableTrimming)
	}
	return sq.q.Select(sortSeries, hints, matchers...)
}

func (sq *storageQuerier) LabelValues(name string, matchers ...*labels.Matcher) ([]string, storage.Warnings, error) {
	_ = level.Debug(sq.log).Log("msg", "labVal", "name", name, "m", matchers)
	return sq.q.LabelValues(name, matchers...)
}

func (sq *storageQuerier) LabelNames(matchers ...*labels.Matcher) ([]string, storage.Warnings, error) {
	_ = level.Debug(sq.log).Log("msg", "labNames", "m", matchers)
	return sq.q.LabelNames(matchers...)
}

func (sq *storageQuerier) Close() error {
	_ = level.Debug(sq.log).Log("msg", "close")
	return sq.q.Close()
}
