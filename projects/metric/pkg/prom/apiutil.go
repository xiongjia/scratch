package prom

import (
	"context"
	"fmt"
	"math"
	"net/http"
	"strconv"
	"time"
	"unsafe"

	"github.com/grafana/regexp"
	jsoniter "github.com/json-iterator/go"
	"github.com/pkg/errors"
	"github.com/prometheus/common/model"
	"github.com/prometheus/prometheus/model/exemplar"
	"github.com/prometheus/prometheus/model/labels"
	"github.com/prometheus/prometheus/model/relabel"

	kitlog "github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/prometheus/prometheus/promql"
	"github.com/prometheus/prometheus/promql/parser"
	"github.com/prometheus/prometheus/storage"
	"github.com/prometheus/prometheus/tsdb"
	"github.com/prometheus/prometheus/util/jsonutil"
	"github.com/prometheus/prometheus/util/stats"
)

type (
	errorType string
	status    string

	apiError struct {
		typ errorType
		err error
	}

	apiFunc func(r *http.Request) apiFuncResult

	apiFuncResult struct {
		data      interface{}
		err       *apiError
		warnings  storage.Warnings
		finalizer func()
	}

	response struct {
		Status    status      `json:"status"`
		Data      interface{} `json:"data,omitempty"`
		ErrorType errorType   `json:"errorType,omitempty"`
		Error     string      `json:"error,omitempty"`
		Warnings  []string    `json:"warnings,omitempty"`
	}

	queryData struct {
		ResultType parser.ValueType `json:"resultType"`
		Result     parser.Value     `json:"result"`
		Stats      stats.QueryStats `json:"stats,omitempty"`
	}

	StatsRenderer func(context.Context, *stats.Statistics, string) stats.QueryStats
)

const (
	errorNone        errorType = ""
	errorTimeout     errorType = "timeout"
	errorCanceled    errorType = "canceled"
	errorExec        errorType = "execution"
	errorBadData     errorType = "bad_data"
	errorInternal    errorType = "internal"
	errorUnavailable errorType = "unavailable"
	errorNotFound    errorType = "not_found"

	statusSuccess status = "success"
	statusError   status = "error"

	// Non-standard status code (originally introduced by nginx) for the case when a client closes
	// the connection while the server is still processing the request.
	statusClientClosedConnection = 499
)

var (
	minTime = time.Unix(math.MinInt64/1000+62135596801, 0).UTC()
	maxTime = time.Unix(math.MaxInt64/1000-62135596801, 999999999).UTC()

	minTimeFormatted = minTime.Format(time.RFC3339Nano)
	maxTimeFormatted = maxTime.Format(time.RFC3339Nano)
)

func init() {
	jsoniter.RegisterTypeEncoderFunc("promql.Series", marshalSeriesJSON, marshalSeriesJSONIsEmpty)
	jsoniter.RegisterTypeEncoderFunc("promql.Sample", marshalSampleJSON, marshalSampleJSONIsEmpty)
	jsoniter.RegisterTypeEncoderFunc("promql.FPoint", marshalFPointJSON, marshalPointJSONIsEmpty)
	jsoniter.RegisterTypeEncoderFunc("promql.HPoint", marshalHPointJSON, marshalPointJSONIsEmpty)
	jsoniter.RegisterTypeEncoderFunc("exemplar.Exemplar", marshalExemplarJSON, marshalExemplarJSONEmpty)
}

func compileCORSRegexString(logger kitlog.Logger, s string) (*regexp.Regexp, error) {
	r, err := relabel.NewRegexp(s)
	if err != nil {
		_ = level.Error(logger).Log("msg", "compile cors reg", "err", err, "reg", s)
		return nil, err
	}
	return r.Regexp, nil
}

func setUnavailStatusOnTSDBNotReady(r apiFuncResult) apiFuncResult {
	if r.err != nil && errors.Cause(r.err.err) == tsdb.ErrNotReady {
		r.err.typ = errorUnavailable
	}
	return r
}

func defaultStatsRenderer(_ context.Context, s *stats.Statistics, param string) stats.QueryStats {
	if param != "" {
		return stats.NewQueryStats(s)
	}
	return nil
}

func returnAPIError(err error) *apiError {
	if err == nil {
		return nil
	}

	cause := errors.Unwrap(err)
	if cause == nil {
		cause = err
	}

	switch cause.(type) {
	case promql.ErrQueryCanceled:
		return &apiError{errorCanceled, err}
	case promql.ErrQueryTimeout:
		return &apiError{errorTimeout, err}
	case promql.ErrStorage:
		return &apiError{errorInternal, err}
	}

	if errors.Is(err, context.Canceled) {
		return &apiError{errorCanceled, err}
	}

	return &apiError{errorExec, err}
}

func invalidParamError(err error, parameter string) apiFuncResult {
	return apiFuncResult{nil, &apiError{
		errorBadData, errors.Wrapf(err, "invalid parameter %q", parameter),
	}, nil, nil}
}

func extractQueryOpts(r *http.Request) (*promql.QueryOpts, error) {
	opts := &promql.QueryOpts{
		EnablePerStepStats: r.FormValue("stats") == "all",
	}
	if strDuration := r.FormValue("lookback_delta"); strDuration != "" {
		duration, err := parseDuration(strDuration)
		if err != nil {
			return nil, fmt.Errorf("error parsing lookback delta duration: %w", err)
		}
		opts.LookbackDelta = duration
	}
	return opts, nil
}

func parseMatchersParam(matchers []string) ([][]*labels.Matcher, error) {
	var matcherSets [][]*labels.Matcher
	for _, s := range matchers {
		matchers, err := parser.ParseMetricSelector(s)
		if err != nil {
			return nil, err
		}
		matcherSets = append(matcherSets, matchers)
	}

OUTER:
	for _, ms := range matcherSets {
		for _, lm := range ms {
			if lm != nil && !lm.Matches("") {
				continue OUTER
			}
		}
		return nil, errors.New("match[] must contain at least one non-empty matcher")
	}
	return matcherSets, nil
}

func parseDuration(s string) (time.Duration, error) {
	if d, err := strconv.ParseFloat(s, 64); err == nil {
		ts := d * float64(time.Second)
		if ts > float64(math.MaxInt64) || ts < float64(math.MinInt64) {
			return 0, errors.Errorf("cannot parse %q to a valid duration. It overflows int64", s)
		}
		return time.Duration(ts), nil
	}
	if d, err := model.ParseDuration(s); err == nil {
		return time.Duration(d), nil
	}
	return 0, errors.Errorf("cannot parse %q to a valid duration", s)
}

func parseTimeParam(r *http.Request, paramName string, defaultValue time.Time) (time.Time, error) {
	val := r.FormValue(paramName)
	if val == "" {
		return defaultValue, nil
	}
	result, err := parseTime(val)
	if err != nil {
		return time.Time{}, errors.Wrapf(err, "Invalid time value for '%s'", paramName)
	}
	return result, nil
}

func parseTime(s string) (time.Time, error) {
	if t, err := strconv.ParseFloat(s, 64); err == nil {
		s, ns := math.Modf(t)
		ns = math.Round(ns*1000) / 1000
		return time.Unix(int64(s), int64(ns*float64(time.Second))).UTC(), nil
	}
	if t, err := time.Parse(time.RFC3339Nano, s); err == nil {
		return t, nil
	}

	// Stdlib's time parser can only handle 4 digit years. As a workaround until
	// that is fixed we want to at least support our own boundary times.
	// Context: https://github.com/prometheus/client_golang/issues/614
	// Upstream issue: https://github.com/golang/go/issues/20555
	switch s {
	case minTimeFormatted:
		return minTime, nil
	case maxTimeFormatted:
		return maxTime, nil
	}
	return time.Time{}, errors.Errorf("cannot parse %q to a valid timestamp", s)
}

// marshalSeriesJSON writes something like the following:
//
//	{
//	   "metric" : {
//	      "__name__" : "up",
//	      "job" : "prometheus",
//	      "instance" : "localhost:9090"
//	   },
//	   "values": [
//	      [ 1435781451.781, "1" ],
//	      < more values>
//	   ],
//	   "histograms": [
//	      [ 1435781451.781, { < histogram, see jsonutil.MarshalHistogram > } ],
//	      < more histograms >
//	   ],
//	},
func marshalSeriesJSON(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	s := *((*promql.Series)(ptr))
	stream.WriteObjectStart()
	stream.WriteObjectField(`metric`)
	m, err := s.Metric.MarshalJSON()
	if err != nil {
		stream.Error = err
		return
	}
	stream.SetBuffer(append(stream.Buffer(), m...))

	for i, p := range s.Floats {
		stream.WriteMore()
		if i == 0 {
			stream.WriteObjectField(`values`)
			stream.WriteArrayStart()
		}
		marshalFPointJSON(unsafe.Pointer(&p), stream)
	}
	if len(s.Floats) > 0 {
		stream.WriteArrayEnd()
	}
	for i, p := range s.Histograms {
		stream.WriteMore()
		if i == 0 {
			stream.WriteObjectField(`histograms`)
			stream.WriteArrayStart()
		}
		marshalHPointJSON(unsafe.Pointer(&p), stream)
	}
	if len(s.Histograms) > 0 {
		stream.WriteArrayEnd()
	}
	stream.WriteObjectEnd()
}

func marshalSeriesJSONIsEmpty(unsafe.Pointer) bool {
	return false
}

// marshalSampleJSON writes something like the following for normal value samples:
//
//	{
//	   "metric" : {
//	      "__name__" : "up",
//	      "job" : "prometheus",
//	      "instance" : "localhost:9090"
//	   },
//	   "value": [ 1435781451.781, "1.234" ]
//	},
//
// For histogram samples, it writes something like this:
//
//	{
//	   "metric" : {
//	      "__name__" : "up",
//	      "job" : "prometheus",
//	      "instance" : "localhost:9090"
//	   },
//	   "histogram": [ 1435781451.781, { < histogram, see jsonutil.MarshalHistogram > } ]
//	},
func marshalSampleJSON(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	s := *((*promql.Sample)(ptr))
	stream.WriteObjectStart()
	stream.WriteObjectField(`metric`)
	m, err := s.Metric.MarshalJSON()
	if err != nil {
		stream.Error = err
		return
	}
	stream.SetBuffer(append(stream.Buffer(), m...))
	stream.WriteMore()
	if s.H == nil {
		stream.WriteObjectField(`value`)
	} else {
		stream.WriteObjectField(`histogram`)
	}
	stream.WriteArrayStart()
	jsonutil.MarshalTimestamp(s.T, stream)
	stream.WriteMore()
	if s.H == nil {
		jsonutil.MarshalFloat(s.F, stream)
	} else {
		jsonutil.MarshalHistogram(s.H, stream)
	}
	stream.WriteArrayEnd()
	stream.WriteObjectEnd()
}

func marshalSampleJSONIsEmpty(unsafe.Pointer) bool {
	return false
}

// marshalFPointJSON writes `[ts, "1.234"]`.
func marshalFPointJSON(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	p := *((*promql.FPoint)(ptr))
	stream.WriteArrayStart()
	jsonutil.MarshalTimestamp(p.T, stream)
	stream.WriteMore()
	jsonutil.MarshalFloat(p.F, stream)
	stream.WriteArrayEnd()
}

// marshalHPointJSON writes `[ts, { < histogram, see jsonutil.MarshalHistogram > } ]`.
func marshalHPointJSON(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	p := *((*promql.HPoint)(ptr))
	stream.WriteArrayStart()
	jsonutil.MarshalTimestamp(p.T, stream)
	stream.WriteMore()
	jsonutil.MarshalHistogram(p.H, stream)
	stream.WriteArrayEnd()
}

func marshalPointJSONIsEmpty(unsafe.Pointer) bool {
	return false
}

// marshalExemplarJSON writes.
//
//	{
//	   labels: <labels>,
//	   value: "<string>",
//	   timestamp: <float>
//	}
func marshalExemplarJSON(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	p := *((*exemplar.Exemplar)(ptr))
	stream.WriteObjectStart()

	// "labels" key.
	stream.WriteObjectField(`labels`)
	lbls, err := p.Labels.MarshalJSON()
	if err != nil {
		stream.Error = err
		return
	}
	stream.SetBuffer(append(stream.Buffer(), lbls...))

	// "value" key.
	stream.WriteMore()
	stream.WriteObjectField(`value`)
	jsonutil.MarshalFloat(p.Value, stream)

	// "timestamp" key.
	stream.WriteMore()
	stream.WriteObjectField(`timestamp`)
	jsonutil.MarshalTimestamp(p.Ts, stream)

	stream.WriteObjectEnd()
}

func marshalExemplarJSONEmpty(unsafe.Pointer) bool {
	return false
}
