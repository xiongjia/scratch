package prom

import (
	kitlog "github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/prometheus/prometheus/model/exemplar"
	"github.com/prometheus/prometheus/model/histogram"
	"github.com/prometheus/prometheus/model/labels"
	"github.com/prometheus/prometheus/model/metadata"
	"github.com/prometheus/prometheus/storage"
	"github.com/prometheus/prometheus/tsdb"
)

type (
	fanoutAppender struct {
		log      kitlog.Logger
		appender storage.Appender
	}

	notReadyAppend struct{}
)

func (notReadyAppend) AppendHistogram(ref storage.SeriesRef, l labels.Labels, t int64, h *histogram.Histogram, fh *histogram.FloatHistogram) (storage.SeriesRef, error) {
	return 0, tsdb.ErrNotReady
}

func (notReadyAppend) AppendHistogramCTZeroSample(ref storage.SeriesRef, l labels.Labels, t, ct int64, h *histogram.Histogram, fh *histogram.FloatHistogram) (storage.SeriesRef, error) {
	return 0, tsdb.ErrNotReady
}

func (notReadyAppend) UpdateMetadata(ref storage.SeriesRef, l labels.Labels, m metadata.Metadata) (storage.SeriesRef, error) {
	return 0, tsdb.ErrNotReady
}

func (notReadyAppend) Append(ref storage.SeriesRef, l labels.Labels, t int64, v float64) (storage.SeriesRef, error) {
	return 0, tsdb.ErrNotReady
}

func (notReadyAppend) Commit() error {
	return tsdb.ErrNotReady
}

func (notReadyAppend) Rollback() error {
	return tsdb.ErrNotReady
}

func (notReadyAppend) AppendExemplar(ref storage.SeriesRef, l labels.Labels, e exemplar.Exemplar) (storage.SeriesRef, error) {
	return 0, tsdb.ErrNotReady
}

func NewFanoutAppender(appender storage.Appender, log kitlog.Logger) storage.Appender {
	return &fanoutAppender{log: log, appender: appender}
}

func (f *fanoutAppender) Append(ref storage.SeriesRef, l labels.Labels, t int64, v float64) (storage.SeriesRef, error) {
	_ = level.Debug(f.log).Log("msg", "append", "ref", ref, "lab", l, "ts", t, "val", v)
	return f.appender.Append(ref, l, t, v)
}

func (f *fanoutAppender) AppendExemplar(ref storage.SeriesRef, l labels.Labels, e exemplar.Exemplar) (storage.SeriesRef, error) {
	_ = level.Debug(f.log).Log("msg", "appendExemplar", "ref", ref, "lab", l, "ex", e)
	return f.appender.AppendExemplar(ref, l, e)
}

func (f *fanoutAppender) AppendHistogram(ref storage.SeriesRef, l labels.Labels, t int64, v1 *histogram.Histogram, v2 *histogram.FloatHistogram) (storage.SeriesRef, error) {
	_ = level.Debug(f.log).Log("msg", "appendHistogram", "ref", ref, "lab", l)
	return f.appender.AppendHistogram(ref, l, t, v1, v2)
}

func (f *fanoutAppender) UpdateMetadata(ref storage.SeriesRef, l labels.Labels, m metadata.Metadata) (storage.SeriesRef, error) {
	_ = level.Debug(f.log).Log("msg", "updateMetadata", "ref", ref, "lab", l, "m", m)
	return f.appender.UpdateMetadata(ref, l, m)
}

func (f *fanoutAppender) Commit() error {
	_ = level.Debug(f.log).Log("msg", "commit")
	return f.appender.Commit()
}

func (f *fanoutAppender) Rollback() error {
	_ = level.Debug(f.log).Log("msg", "rollback")
	return f.appender.Commit()
}
