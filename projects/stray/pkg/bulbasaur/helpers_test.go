package bulbasaur_test

import (
	"context"

	"github.com/prometheus/prometheus/model/exemplar"
	"github.com/prometheus/prometheus/model/histogram"
	"github.com/prometheus/prometheus/model/labels"
	"github.com/prometheus/prometheus/model/metadata"
	"github.com/prometheus/prometheus/storage"
)

type nopAppendable struct{}

func (a nopAppendable) Appender(_ context.Context) storage.Appender {
	return nopAppender{}
}

type nopAppender struct{}

func (a nopAppender) Append(storage.SeriesRef, labels.Labels, int64, float64) (storage.SeriesRef, error) {
	return 0, nil
}

func (a nopAppender) AppendExemplar(storage.SeriesRef, labels.Labels, exemplar.Exemplar) (storage.SeriesRef, error) {
	return 0, nil
}

func (a nopAppender) AppendHistogram(storage.SeriesRef, labels.Labels, int64, *histogram.Histogram, *histogram.FloatHistogram) (storage.SeriesRef, error) {
	return 0, nil
}

func (a nopAppender) UpdateMetadata(storage.SeriesRef, labels.Labels, metadata.Metadata) (storage.SeriesRef, error) {
	return 0, nil
}

func (a nopAppender) AppendCTZeroSample(storage.SeriesRef, labels.Labels, int64, int64) (storage.SeriesRef, error) {
	return 0, nil
}

func (a nopAppender) Commit() error {
	return nil
}

func (a nopAppender) Rollback() error {
	return nil
}
