package prom

import (
	"context"

	kitlog "github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/prometheus/prometheus/model/labels"
	"github.com/prometheus/prometheus/storage"
)

type (
	dbQuerier struct {
		log     kitlog.Logger
		ctx     context.Context
		storage *storageDb

		mint int64
		maxt int64
	}

	dbResultSeriesSet struct {
		log kitlog.Logger
		ctx context.Context
		itr int

		err  error
		data *rcdSet
	}
)

func (s *dbResultSeriesSet) Next() bool {
	if s.err != nil {
		return false
	}

	if s.data == nil {
		return false
	}

	if s.itr >= len(s.data.items) {
		return false
	}
	s.itr++
	return true

}

func (s *dbResultSeriesSet) At() storage.Series {
	return makeSeries(s.data, s.itr-1)
}

func (s *dbResultSeriesSet) Err() error {
	return s.err
}

func (s *dbResultSeriesSet) Warnings() storage.Warnings {
	return nil
}

func (q *dbQuerier) Select(sortSeries bool, hints *storage.SelectHints, matchers ...*labels.Matcher) storage.SeriesSet {
	_ = level.Debug(q.log).Log("msg", "Select", "sortSeries", sortSeries, "hints", hints, "matchers", matchers)

	result := &dbResultSeriesSet{err: nil, log: q.log, ctx: q.ctx, itr: 0}
	if hints == nil {
		hints = &storage.SelectHints{
			Start: q.mint,
			End:   q.mint,
		}
	}
	labs, err := q.storage.dbFindMatchLabs(matchers...)
	if err != nil {
		result.err = err
		return result
	}

	rcdSet, err := q.storage.dbFindRcdSet(labs, hints)
	if err != nil {
		result.err = err
		return result
	}
	result.data = rcdSet
	return result
}

func (q *dbQuerier) LabelValues(name string, matchers ...*labels.Matcher) ([]string, storage.Warnings, error) {
	_ = level.Debug(q.log).Log("msg", "LabelValues", "name", name, "matchers", matchers)
	labs, err := q.storage.dbFindMatchLabs(matchers...)
	if err != nil {
		return []string{}, []error{err}, nil
	}

	result := make([]string, 0)
	for _, lab := range labs {
		for _, itr := range lab.Lab {
			if itr.Name != name {
				continue
			}
			result = append(result, itr.Value)
		}
	}
	return result, []error{}, nil

}

func (q *dbQuerier) LabelNames(matchers ...*labels.Matcher) ([]string, storage.Warnings, error) {
	_ = level.Debug(q.log).Log("msg", "LabelNames", "matchers", matchers)
	labs, err := q.storage.dbFindMatchLabs(matchers...)
	if err != nil {
		return []string{}, []error{err}, nil
	}

	result := make([]string, 0)
	for _, lab := range labs {
		for _, itr := range lab.Lab {
			result = append(result, itr.Name)
		}
	}
	return result, []error{}, nil
}

func (q *dbQuerier) Close() error {
	return nil
}
