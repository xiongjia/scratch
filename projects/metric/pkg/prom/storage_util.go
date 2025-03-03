package prom

import (
	"encoding/json"
	"slices"
	"strings"
	"time"

	kitlog "github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/prometheus/prometheus/model/labels"
	"github.com/prometheus/prometheus/model/metadata"
	"github.com/prometheus/prometheus/storage"
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

func labsSort(src labels.Labels) labels.Labels {
	slices.SortFunc(src, func(a labels.Label, b labels.Label) int {
		return strings.Compare(a.Name, b.Name)
	})
	return src
}

func labMetadataToJson(src *metadata.Metadata) (string, error) {
	if src == nil {
		return "", nil
	}
	b, err := json.Marshal(src)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func labMetadataFromJsonDefault(src string, defaultVal *metadata.Metadata) *metadata.Metadata {
	m, err := labMetadataFromJson(src)
	if err != nil {
		return defaultVal
	}
	return m
}

func labMetadataFromJson(src string) (*metadata.Metadata, error) {
	if len(src) == 0 {
		return nil, nil
	}

	var m metadata.Metadata
	if err := json.Unmarshal([]byte(src), &m); err != nil {
		return nil, err
	}
	return &m, nil
}

func labMetadataEqual(m1 *metadata.Metadata, m2 *metadata.Metadata) bool {
	if m1 == nil && m2 == nil {
		return true
	}
	if m1 == nil && m2 != nil {
		return false
	}
	if m1 != nil && m2 == nil {
		return false
	}
	if (m1.Help == m2.Help) && (m1.Unit == m2.Unit) && (m1.Type == m2.Type) {
		return true
	}
	return false
}

func labMetadataCopy(src *metadata.Metadata) *metadata.Metadata {
	if src == nil {
		return nil
	}
	return &metadata.Metadata{
		Type: src.Type,
		Unit: strings.Clone(src.Unit),
		Help: strings.Clone(src.Help),
	}
}

func labsToJson(src labels.Labels) (string, error) {
	b, err := src.MarshalJSON()
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func labsFromJson(src string) (labels.Labels, error) {
	var m map[string]string
	if err := json.Unmarshal([]byte(src), &m); err != nil {
		return nil, err
	}
	return labels.FromMap(m), nil
}

func labMatchItem(lab *labels.Label, matcher *labels.Matcher) bool {
	if matcher == nil || lab == nil {
		return false
	}
	if matcher.Name != lab.Name {
		return false
	}
	return matcher.Matches(lab.Value)
}

func labsMatch(labs labels.Labels, matchers ...*labels.Matcher) bool {
	if len(labs) == 0 {
		return false
	}
	if len(matchers) == 0 {
		return true
	}
	for _, m := range matchers {
		found := false
		for _, lab := range labs {
			if labMatchItem(&lab, m) {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}
