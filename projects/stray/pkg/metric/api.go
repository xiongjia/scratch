package metric

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"math"
	"net/http"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/grafana/regexp"
	"github.com/munnerz/goautoneg"
	"github.com/prometheus/common/model"
	"github.com/prometheus/common/route"
	"github.com/prometheus/prometheus/model/labels"
	"github.com/prometheus/prometheus/model/timestamp"
	"github.com/prometheus/prometheus/promql"
	"github.com/prometheus/prometheus/promql/parser"
	"github.com/prometheus/prometheus/storage"
	"github.com/prometheus/prometheus/tsdb"
	"github.com/prometheus/prometheus/util/annotations"
	"github.com/prometheus/prometheus/util/httputil"
	"github.com/prometheus/prometheus/util/stats"

	jsoniter "github.com/json-iterator/go"
)

type (
	StatsRenderer func(context.Context, *stats.Statistics, string) stats.QueryStats

	EngineApi struct {
		promQuerier *PromQuerier
		Queryable   *PromStorage

		mux           *http.ServeMux
		httpHandler   http.Handler
		cors          *regexp.Regexp
		logger        *slog.Logger
		codecs        []Codec
		statsRenderer StatsRenderer
		ready         func(http.HandlerFunc) http.HandlerFunc
		now           func() time.Time
	}

	EngineApiOptions struct {
		ServeMux    *http.ServeMux
		HttpHandler http.Handler
		Querier     *PromQuerier
		Storage     *PromStorage
	}

	errorType string
	status    string

	apiError struct {
		typ errorType
		err error
	}

	apiFuncResult struct {
		data      interface{}
		err       *apiError
		warnings  annotations.Annotations
		finalizer func()
	}

	apiFunc func(r *http.Request) apiFuncResult

	Response struct {
		Status    status      `json:"status"`
		Data      interface{} `json:"data,omitempty"`
		ErrorType errorType   `json:"errorType,omitempty"`
		Error     string      `json:"error,omitempty"`
		Warnings  []string    `json:"warnings,omitempty"`
		Infos     []string    `json:"infos,omitempty"`
	}

	QueryData struct {
		ResultType parser.ValueType `json:"resultType"`
		Result     parser.Value     `json:"result"`
		Stats      stats.QueryStats `json:"stats,omitempty"`
	}
)

const (
	errorNone          errorType = ""
	errorTimeout       errorType = "timeout"
	errorCanceled      errorType = "canceled"
	errorExec          errorType = "execution"
	errorBadData       errorType = "bad_data"
	errorInternal      errorType = "internal"
	errorUnavailable   errorType = "unavailable"
	errorNotFound      errorType = "not_found"
	errorNotAcceptable errorType = "not_acceptable"

	statusSuccess status = "success"
	statusError   status = "error"

	// Non-standard status code (originally introduced by nginx) for the case when a client closes
	// the connection while the server is still processing the request.
	statusClientClosedConnection = 499
)

var (
	// MinTime is the default timestamp used for the start of optional time ranges.
	// Exposed to let downstream projects reference it.
	//
	// Historical note: This should just be time.Unix(math.MinInt64/1000, 0).UTC(),
	// but it was set to a higher value in the past due to a misunderstanding.
	// The value is still low enough for practical purposes, so we don't want
	// to change it now, avoiding confusion for importers of this variable.
	MinTime = time.Unix(math.MinInt64/1000+62135596801, 0).UTC()

	// MaxTime is the default timestamp used for the end of optional time ranges.
	// Exposed to let downstream projects to reference it.
	//
	// Historical note: This should just be time.Unix(math.MaxInt64/1000, 0).UTC(),
	// but it was set to a lower value in the past due to a misunderstanding.
	// The value is still high enough for practical purposes, so we don't want
	// to change it now, avoiding confusion for importers of this variable.
	MaxTime = time.Unix(math.MaxInt64/1000-62135596801, 999999999).UTC()

	minTimeFormatted = MinTime.Format(time.RFC3339Nano)
	maxTimeFormatted = MaxTime.Format(time.RFC3339Nano)
)

func parseTimeParam(r *http.Request, paramName string, defaultValue time.Time) (time.Time, error) {
	val := r.FormValue(paramName)
	if val == "" {
		return defaultValue, nil
	}
	result, err := parseTime(val)
	if err != nil {
		return time.Time{}, fmt.Errorf("Invalid time value for '%s': %w", paramName, err)
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
		return MinTime, nil
	case maxTimeFormatted:
		return MaxTime, nil
	}
	return time.Time{}, fmt.Errorf("cannot parse %q to a valid timestamp", s)
}

func invalidParamError(err error, parameter string) apiFuncResult {
	return apiFuncResult{nil, &apiError{
		errorBadData, fmt.Errorf("invalid parameter %q: %w", parameter, err),
	}, nil, nil}
}

// parseLimitParam returning 0 means no limit is to be applied.
func parseLimitParam(limitStr string) (limit int, err error) {
	if limitStr == "" {
		return limit, nil
	}

	limit, err = strconv.Atoi(limitStr)
	if err != nil {
		return limit, err
	}
	if limit < 0 {
		return limit, errors.New("limit must be non-negative")
	}

	return limit, nil
}

// toHintLimit increases the API limit, as returned by parseLimitParam, by 1.
// This allows for emitting warnings when the results are truncated.
func toHintLimit(limit int) int {
	// 0 means no limit and avoid int overflow
	if limit > 0 && limit < math.MaxInt {
		return limit + 1
	}
	return limit
}

func parseDuration(s string) (time.Duration, error) {
	if d, err := strconv.ParseFloat(s, 64); err == nil {
		ts := d * float64(time.Second)
		if ts > float64(math.MaxInt64) || ts < float64(math.MinInt64) {
			return 0, fmt.Errorf("cannot parse %q to a valid duration. It overflows int64", s)
		}
		return time.Duration(ts), nil
	}
	if d, err := model.ParseDuration(s); err == nil {
		return time.Duration(d), nil
	}
	return 0, fmt.Errorf("cannot parse %q to a valid duration", s)
}

func extractQueryOpts(r *http.Request) (promql.QueryOpts, error) {
	var duration time.Duration

	if strDuration := r.FormValue("lookback_delta"); strDuration != "" {
		parsedDuration, err := parseDuration(strDuration)
		if err != nil {
			return nil, fmt.Errorf("error parsing lookback delta duration: %w", err)
		}
		duration = parsedDuration
	}

	return promql.NewPrometheusQueryOpts(r.FormValue("stats") == "all", duration), nil
}

func setUnavailStatusOnTSDBNotReady(r apiFuncResult) apiFuncResult {
	if r.err != nil && errors.Is(r.err.err, tsdb.ErrNotReady) {
		r.err.typ = errorUnavailable
	}
	return r
}

func returnAPIError(err error) *apiError {
	if err == nil {
		return nil
	}

	var eqc promql.ErrQueryCanceled
	var eqt promql.ErrQueryTimeout
	var es promql.ErrStorage
	switch {
	case errors.As(err, &eqc):
		return &apiError{errorCanceled, err}
	case errors.As(err, &eqt):
		return &apiError{errorTimeout, err}
	case errors.As(err, &es):
		return &apiError{errorInternal, err}
	}

	if errors.Is(err, context.Canceled) {
		return &apiError{errorCanceled, err}
	}

	return &apiError{errorExec, err}
}

func DefaultStatsRenderer(_ context.Context, s *stats.Statistics, param string) stats.QueryStats {
	if param != "" {
		return stats.NewQueryStats(s)
	}
	return nil
}

func (api *EngineApi) respondError(w http.ResponseWriter, apiErr *apiError, data interface{}) {
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	b, err := json.Marshal(&Response{
		Status:    statusError,
		ErrorType: apiErr.typ,
		Error:     apiErr.err.Error(),
		Data:      data,
	})
	if err != nil {
		api.logger.Error("error marshaling json response", "err", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var code int
	switch apiErr.typ {
	case errorBadData:
		code = http.StatusBadRequest
	case errorExec:
		code = http.StatusUnprocessableEntity
	case errorCanceled:
		code = statusClientClosedConnection
	case errorTimeout:
		code = http.StatusServiceUnavailable
	case errorInternal:
		code = http.StatusInternalServerError
	case errorNotFound:
		code = http.StatusNotFound
	case errorNotAcceptable:
		code = http.StatusNotAcceptable
	default:
		code = http.StatusInternalServerError
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if n, err := w.Write(b); err != nil {
		api.logger.Error("error writing response", "bytesWritten", n, "err", err)
	}
}

func (api *EngineApi) negotiateCodec(req *http.Request, resp *Response) (Codec, error) {
	for _, clause := range goautoneg.ParseAccept(req.Header.Get("Accept")) {
		for _, codec := range api.codecs {
			if codec.ContentType().Satisfies(clause) && codec.CanEncode(resp) {
				return codec, nil
			}
		}
	}

	defaultCodec := api.codecs[0]
	if !defaultCodec.CanEncode(resp) {
		return nil, fmt.Errorf("cannot encode response as %s", defaultCodec.ContentType())
	}

	return defaultCodec, nil
}

func parseMatchersParam(matchers []string) ([][]*labels.Matcher, error) {
	matcherSets, err := parser.ParseMetricSelectors(matchers)
	if err != nil {
		return nil, err
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

func (api *EngineApi) respond(w http.ResponseWriter, req *http.Request, data interface{}, warnings annotations.Annotations, query string) {
	statusMessage := statusSuccess
	warn, info := warnings.AsStrings(query, 10, 10)

	resp := &Response{
		Status:   statusMessage,
		Data:     data,
		Warnings: warn,
		Infos:    info,
	}

	codec, err := api.negotiateCodec(req, resp)
	if err != nil {
		api.respondError(w, &apiError{errorNotAcceptable, err}, nil)
		return
	}

	b, err := codec.Encode(resp)
	if err != nil {
		api.logger.Error("error marshaling response", "url", req.URL, "err", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", codec.ContentType().String())
	w.WriteHeader(http.StatusOK)
	if n, err := w.Write(b); err != nil {
		api.logger.Error("error writing response", "url", req.URL, "bytesWritten", n, "err", err)
	}
}

func (api *EngineApi) Register() {
	wrap := func(f apiFunc) http.HandlerFunc {
		hf := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			httputil.SetCORS(w, api.cors, r)
			result := setUnavailStatusOnTSDBNotReady(f(r))
			if result.finalizer != nil {
				defer result.finalizer()
			}
			if result.err != nil {
				api.respondError(w, result.err, result.data)
				return
			}

			if result.data != nil {
				api.respond(w, r, result.data, result.warnings, r.FormValue("query"))
				return
			}
			w.WriteHeader(http.StatusNoContent)
		})
		return httputil.CompressionHandler{Handler: hf}.ServeHTTP
	}

	apiRoute := route.New()
	apiRoute.Get("/mon/api/v1/query", wrap(api.query))
	apiRoute.Get("/labels", wrap(api.labelNames))
	apiRoute.Get("/label/:name/values", wrap(api.labelValues))

	api.mux.Handle("/", apiRoute)
}

func (api *EngineApi) labelValues(r *http.Request) (result apiFuncResult) {
	ctx := r.Context()
	name := route.Param(ctx, "name")

	if strings.HasPrefix(name, "U__") {
		name = model.UnescapeName(name, model.ValueEncodingEscaping)
	}

	label := model.LabelName(name)
	if !label.IsValid() {
		return apiFuncResult{nil, &apiError{errorBadData, fmt.Errorf("invalid label name: %q", name)}, nil, nil}
	}

	limit, err := parseLimitParam(r.FormValue("limit"))
	if err != nil {
		return invalidParamError(err, "limit")
	}

	start, err := parseTimeParam(r, "start", MinTime)
	if err != nil {
		return invalidParamError(err, "start")
	}
	end, err := parseTimeParam(r, "end", MaxTime)
	if err != nil {
		return invalidParamError(err, "end")
	}

	matcherSets, err := parseMatchersParam(r.Form["match[]"])
	if err != nil {
		return apiFuncResult{nil, &apiError{errorBadData, err}, nil, nil}
	}

	hints := &storage.LabelHints{
		Limit: toHintLimit(limit),
	}

	q, err := api.Queryable.Querier(timestamp.FromTime(start), timestamp.FromTime(end))
	if err != nil {
		return apiFuncResult{nil, &apiError{errorExec, err}, nil, nil}
	}
	// From now on, we must only return with a finalizer in the result (to
	// be called by the caller) or call q.Close ourselves (which is required
	// in the case of a panic).
	defer func() {
		if result.finalizer == nil {
			q.Close()
		}
	}()
	closer := func() {
		q.Close()
	}

	var (
		vals     []string
		warnings annotations.Annotations
	)
	if len(matcherSets) > 1 {
		var callWarnings annotations.Annotations
		labelValuesSet := make(map[string]struct{})
		for _, matchers := range matcherSets {
			vals, callWarnings, err = q.LabelValues(ctx, name, hints, matchers...)
			if err != nil {
				return apiFuncResult{nil, &apiError{errorExec, err}, warnings, closer}
			}
			warnings.Merge(callWarnings)
			for _, val := range vals {
				labelValuesSet[val] = struct{}{}
			}
		}

		vals = make([]string, 0, len(labelValuesSet))
		for val := range labelValuesSet {
			vals = append(vals, val)
		}
	} else {
		var matchers []*labels.Matcher
		if len(matcherSets) == 1 {
			matchers = matcherSets[0]
		}
		vals, warnings, err = q.LabelValues(ctx, name, hints, matchers...)
		if err != nil {
			return apiFuncResult{nil, &apiError{errorExec, err}, warnings, closer}
		}

		if vals == nil {
			vals = []string{}
		}
	}

	slices.Sort(vals)

	if limit > 0 && len(vals) > limit {
		vals = vals[:limit]
		warnings = warnings.Add(errors.New("results truncated due to limit"))
	}

	return apiFuncResult{vals, nil, warnings, closer}
}

func (api *EngineApi) labelNames(r *http.Request) apiFuncResult {
	limit, err := parseLimitParam(r.FormValue("limit"))
	if err != nil {
		return invalidParamError(err, "limit")
	}

	start, err := parseTimeParam(r, "start", MinTime)
	if err != nil {
		return invalidParamError(err, "start")
	}
	end, err := parseTimeParam(r, "end", MaxTime)
	if err != nil {
		return invalidParamError(err, "end")
	}

	matcherSets, err := parseMatchersParam(r.Form["match[]"])
	if err != nil {
		return apiFuncResult{nil, &apiError{errorBadData, err}, nil, nil}
	}

	hints := &storage.LabelHints{
		Limit: toHintLimit(limit),
	}

	q, err := api.Queryable.Querier(timestamp.FromTime(start), timestamp.FromTime(end))
	if err != nil {
		return apiFuncResult{nil, returnAPIError(err), nil, nil}
	}
	defer q.Close()

	var (
		names    []string
		warnings annotations.Annotations
	)
	if len(matcherSets) > 1 {
		labelNamesSet := make(map[string]struct{})

		for _, matchers := range matcherSets {
			vals, callWarnings, err := q.LabelNames(r.Context(), hints, matchers...)
			if err != nil {
				return apiFuncResult{nil, returnAPIError(err), warnings, nil}
			}

			warnings.Merge(callWarnings)
			for _, val := range vals {
				labelNamesSet[val] = struct{}{}
			}
		}

		// Convert the map to an array.
		names = make([]string, 0, len(labelNamesSet))
		for key := range labelNamesSet {
			names = append(names, key)
		}
		slices.Sort(names)
	} else {
		var matchers []*labels.Matcher
		if len(matcherSets) == 1 {
			matchers = matcherSets[0]
		}
		names, warnings, err = q.LabelNames(r.Context(), hints, matchers...)
		if err != nil {
			return apiFuncResult{nil, &apiError{errorExec, err}, warnings, nil}
		}
	}

	if names == nil {
		names = []string{}
	}

	if limit > 0 && len(names) > limit {
		names = names[:limit]
		warnings = warnings.Add(errors.New("results truncated due to limit"))
	}
	return apiFuncResult{names, nil, warnings, nil}
}

func (api *EngineApi) query(r *http.Request) (result apiFuncResult) {
	ts, err := parseTimeParam(r, "time", api.now())
	if err != nil {
		return invalidParamError(err, "time")
	}
	ctx := r.Context()
	if to := r.FormValue("timeout"); to != "" {
		var cancel context.CancelFunc
		timeout, err := parseDuration(to)
		if err != nil {
			return invalidParamError(err, "timeout")
		}

		ctx, cancel = context.WithDeadline(ctx, api.now().Add(timeout))
		defer cancel()
	}

	opts, err := extractQueryOpts(r)
	if err != nil {
		return apiFuncResult{nil, &apiError{errorBadData, err}, nil, nil}
	}
	qry, err := api.promQuerier.QueryEngine.NewInstantQuery(ctx, api.Queryable, opts, r.FormValue("query"), ts)
	if err != nil {
		return invalidParamError(err, "query")
	}

	// From now on, we must only return with a finalizer in the result (to
	// be called by the caller) or call qry.Close ourselves (which is
	// required in the case of a panic).
	defer func() {
		if result.finalizer == nil {
			qry.Close()
		}
	}()

	ctx = httputil.ContextFromRequest(ctx, r)

	res := qry.Exec(ctx)
	if res.Err != nil {
		return apiFuncResult{nil, returnAPIError(res.Err), res.Warnings, qry.Close}
	}

	// Optional stats field in response if parameter "stats" is not empty.
	sr := api.statsRenderer
	if sr == nil {
		sr = DefaultStatsRenderer
	}
	qs := sr(ctx, qry.Stats(), r.FormValue("stats"))

	return apiFuncResult{&QueryData{
		ResultType: res.Value.Type(),
		Result:     res.Value,
		Stats:      qs,
	}, nil, res.Warnings, qry.Close}
}

func NewEngineApi(opts EngineApiOptions) *EngineApi {
	cors, _ := compileCORSRegexString(".*")
	engApi := &EngineApi{
		cors:          cors,
		httpHandler:   opts.HttpHandler,
		mux:           opts.ServeMux,
		promQuerier:   opts.Querier,
		Queryable:     opts.Storage,
		statsRenderer: DefaultStatsRenderer,
		now:           time.Now,
	}

	engApi.codecs = append(engApi.codecs, JSONCodec{})
	engApi.Register()
	return engApi
}
