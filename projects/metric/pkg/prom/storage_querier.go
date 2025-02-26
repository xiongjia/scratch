package prom

import (
	"github.com/prometheus/prometheus/model/labels"
	"github.com/prometheus/prometheus/storage"
)

type (
	storageQuerier struct {
		q storage.Querier
	}
)

func makeStorageQuerier(q storage.Querier) storage.Querier {
	return &storageQuerier{q: q}
}

func (sq *storageQuerier) Select(sortSeries bool, hints *storage.SelectHints, matchers ...*labels.Matcher) storage.SeriesSet {
	return sq.q.Select(sortSeries, hints, matchers...)
}

func (sq *storageQuerier) LabelValues(name string, matchers ...*labels.Matcher) ([]string, storage.Warnings, error) {
	return sq.q.LabelValues(name, matchers...)
}

func (sq *storageQuerier) LabelNames(matchers ...*labels.Matcher) ([]string, storage.Warnings, error) {
	return sq.q.LabelNames(matchers...)
}

func (sq *storageQuerier) Close() error {
	return sq.q.Close()
}
