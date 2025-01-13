package metric

import (
	"context"
	"log/slog"

	"github.com/grafana/regexp"
	"github.com/prometheus/common/model"
	"github.com/prometheus/prometheus/discovery"
	"github.com/prometheus/prometheus/discovery/targetgroup"
	"github.com/prometheus/prometheus/model/exemplar"
	"github.com/prometheus/prometheus/model/histogram"
	"github.com/prometheus/prometheus/model/labels"
	"github.com/prometheus/prometheus/model/metadata"
	"github.com/prometheus/prometheus/model/relabel"
	"github.com/prometheus/prometheus/storage"
	"github.com/prometheus/prometheus/tsdb"
	"github.com/samber/lo"
)

type (
	StaticTargetGroup struct {
		Source    string
		Addresses []string
	}

	nopAppendable struct{}

	nopAppend struct{}

	notReadyAppend struct{}
)

func NewNopAppend() storage.Appendable {
	return &nopAppendable{}
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

func (notReadyAppend) SetOptions(opts *storage.AppendOptions) {
}

func (notReadyAppend) AppendExemplar(ref storage.SeriesRef, l labels.Labels, e exemplar.Exemplar) (storage.SeriesRef, error) {
	return 0, tsdb.ErrNotReady
}

func (notReadyAppend) AppendHistogram(ref storage.SeriesRef, l labels.Labels, t int64, h *histogram.Histogram, fh *histogram.FloatHistogram) (storage.SeriesRef, error) {
	return 0, tsdb.ErrNotReady
}

func (notReadyAppend) AppendHistogramCTZeroSample(ref storage.SeriesRef, l labels.Labels, t, ct int64, h *histogram.Histogram, fh *histogram.FloatHistogram) (storage.SeriesRef, error) {
	return 0, tsdb.ErrNotReady
}

func (notReadyAppend) UpdateMetadata(ref storage.SeriesRef, l labels.Labels, m metadata.Metadata) (storage.SeriesRef, error) {
	return 0, tsdb.ErrNotReady
}

func (notReadyAppend) AppendCTZeroSample(ref storage.SeriesRef, l labels.Labels, t, ct int64) (storage.SeriesRef, error) {
	return 0, tsdb.ErrNotReady
}

func (nopAppendable) Appender(_ context.Context) storage.Appender {
	return nopAppend{}
}

func (nopAppend) SetOptions(opts *storage.AppendOptions) {}

func (nopAppend) Append(_ storage.SeriesRef, l labels.Labels, t int64, v float64) (storage.SeriesRef, error) {
	slog.Debug("nop append", slog.Any("lab", l), slog.Int64("time", t), slog.Float64("val", v))
	return 0, nil
}

func (nopAppend) AppendExemplar(storage.SeriesRef, labels.Labels, exemplar.Exemplar) (storage.SeriesRef, error) {
	return 0, nil
}

func (nopAppend) AppendHistogram(storage.SeriesRef, labels.Labels, int64, *histogram.Histogram, *histogram.FloatHistogram) (storage.SeriesRef, error) {
	return 0, nil
}

func (nopAppend) AppendHistogramCTZeroSample(ref storage.SeriesRef, l labels.Labels, t, ct int64, h *histogram.Histogram, fh *histogram.FloatHistogram) (storage.SeriesRef, error) {
	return 0, nil
}

func (nopAppend) UpdateMetadata(storage.SeriesRef, labels.Labels, metadata.Metadata) (storage.SeriesRef, error) {
	return 0, nil
}

func (nopAppend) AppendCTZeroSample(storage.SeriesRef, labels.Labels, int64, int64) (storage.SeriesRef, error) {
	return 0, nil
}

func (nopAppend) Commit() error { return nil }

func (nopAppend) Rollback() error { return nil }

func makeComponentLog(log *slog.Logger, component string) *slog.Logger {
	return lo.If(log == nil, slog.Default()).Else(log).With(LOG_COMPONENT_KEY,
		lo.If(component == "", "default").Else(component))
}

func makeDiscoveryStaticConfig(groups ...StaticTargetGroup) discovery.StaticConfig {
	var cfg discovery.StaticConfig
	lo.ForEach(groups, func(grp StaticTargetGroup, _ int) {
		lo.ForEach(grp.Addresses, func(addr string, _ int) {
			cfg = append(cfg, &targetgroup.Group{
				Source: grp.Source,
				Targets: []model.LabelSet{
					{model.AddressLabel: model.LabelValue(addr)},
				},
			})
		})
	})
	return cfg
}

func compileCORSRegexString(s string) (*regexp.Regexp, error) {
	r, err := relabel.NewRegexp(s)
	if err != nil {
		slog.Error("compile cors reg", slog.Any("err", err), slog.String("src", s))
		return nil, err
	}
	return r.Regexp, nil
}
