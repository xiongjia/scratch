package prom

import (
	"context"
	"net/http"
	"strconv"
	"time"

	kitlog "github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/grafana/regexp"
	jsoniter "github.com/json-iterator/go"
	"github.com/pkg/errors"
	"github.com/prometheus/common/model"
	"github.com/prometheus/common/route"
	"github.com/prometheus/prometheus/model/labels"
	"github.com/prometheus/prometheus/model/textparse"
	"github.com/prometheus/prometheus/model/timestamp"
	"github.com/prometheus/prometheus/promql"
	"github.com/prometheus/prometheus/promql/parser"
	"github.com/prometheus/prometheus/scrape"
	"github.com/prometheus/prometheus/storage"
	"github.com/prometheus/prometheus/util/httputil"
	"golang.org/x/exp/slices"
)

type (
	QueryEngine interface {
		SetQueryLogger(l promql.QueryLogger)
		NewInstantQuery(ctx context.Context, q storage.Queryable, opts *promql.QueryOpts, qs string, ts time.Time) (promql.Query, error)
		NewRangeQuery(ctx context.Context, q storage.Queryable, opts *promql.QueryOpts, qs string, start, end time.Time, interval time.Duration) (promql.Query, error)
	}

	TargetRetriever interface {
		TargetsActive() map[string][]*scrape.Target
		TargetsDropped() map[string][]*scrape.Target
	}

	apiMetadata struct {
		Type textparse.MetricType `json:"type"`
		Help string               `json:"help"`
		Unit string               `json:"unit"`
	}

	EngineAPI struct {
		promQuerier *PromQuerier
		promStorage *PromStorage
		promScrape  *PromScrape

		queryable         storage.SampleAndChunkQueryable
		exemplarQueryable storage.ExemplarQueryable
		queryEngine       QueryEngine

		logger          kitlog.Logger
		corsOrigin      *regexp.Regexp
		statsRenderer   StatsRenderer
		now             func() time.Time
		ready           func(http.HandlerFunc) http.HandlerFunc
		targetRetriever func(context.Context) TargetRetriever
	}

	EngineAPIOpts struct {
		Logger  kitlog.Logger
		Querier *PromQuerier
		Storage *PromStorage
		Scrape  *PromScrape
	}
)

func NewEngineAPI(opts EngineAPIOpts) (*EngineAPI, error) {
	cors, err := compileCORSRegexString(opts.Logger, ".*")
	if err != nil {
		_ = level.Error(opts.Logger).Log("msg", "compile cors regex", "err", err)
		return nil, err
	}
	factoryTr := func(_ context.Context) TargetRetriever {
		return opts.Scrape.scrapeMgr
	}
	api := &EngineAPI{
		promQuerier: opts.Querier,
		promStorage: opts.Storage,
		promScrape:  opts.Scrape,

		queryable:         opts.Storage,
		exemplarQueryable: opts.Storage,
		queryEngine:       opts.Querier,

		logger:          kitlog.With(opts.Logger, LOG_COMPONENT_KEY, COMPONENT_API),
		corsOrigin:      cors,
		statsRenderer:   defaultStatsRenderer,
		now:             time.Now,
		ready:           apiReady,
		targetRetriever: factoryTr,
	}
	return api, nil
}

func apiReady(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		f(w, r)
	}
}

func (api *EngineAPI) Register(prefix string) (http.Handler, error) {
	_ = level.Error(api.logger).Log("msg", "binding monitor api handlers", "prefix", prefix)

	wrap := func(f apiFunc) http.HandlerFunc {
		hf := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			httputil.SetCORS(w, api.corsOrigin, r)
			result := setUnavailStatusOnTSDBNotReady(f(r))
			if result.finalizer != nil {
				defer result.finalizer()
			}
			if result.err != nil {
				api.respondError(w, result.err, result.data)
				return
			}

			if result.data != nil {
				api.respond(w, result.data, result.warnings)
				return
			}
			w.WriteHeader(http.StatusNoContent)
		})
		return api.ready(httputil.CompressionHandler{Handler: hf}.ServeHTTP)
	}

	if prefix == "/" {
		prefix = ""
	}
	apiPath := prefix + "/api/v1"

	apiRoute := route.New()
	apiRoute.Get(apiPath+"/query", wrap(api.query))
	apiRoute.Post(apiPath+"/query", wrap(api.query))
	apiRoute.Get(apiPath+"/query_range", wrap(api.queryRange))
	apiRoute.Post(apiPath+"/query_range", wrap(api.queryRange))
	apiRoute.Get(apiPath+"/query_exemplars", wrap(api.queryExemplars))
	apiRoute.Post(apiPath+"/query_exemplars", wrap(api.queryExemplars))
	apiRoute.Get(apiPath+"/series", wrap(api.series))
	apiRoute.Post(apiPath+"/series", wrap(api.series))
	apiRoute.Get(apiPath+"/labels", wrap(api.labelNames))
	apiRoute.Post(apiPath+"/labels", wrap(api.labelNames))
	apiRoute.Get(apiPath+"/label/:name/values", wrap(api.labelValues))
	apiRoute.Post(apiPath+"/label/:name/values", wrap(api.labelValues))
	apiRoute.Get(apiPath+"/metadata", wrap(api.metricMetadata))
	apiRoute.Post(apiPath+"/metadata", wrap(api.metricMetadata))
	apiRoute.Get(apiPath+"/format_query", wrap(api.formatQuery))
	apiRoute.Post(apiPath+"/format_query", wrap(api.formatQuery))
	mux := http.NewServeMux()
	mux.Handle("/", apiRoute)
	return mux, nil
}

func (api *EngineAPI) formatQuery(r *http.Request) (result apiFuncResult) {
	expr, err := parser.ParseExpr(r.FormValue("query"))
	if err != nil {
		return invalidParamError(err, "query")
	}

	return apiFuncResult{expr.Pretty(0), nil, nil, nil}
}

func (api *EngineAPI) queryExemplars(r *http.Request) apiFuncResult {
	start, err := parseTimeParam(r, "start", minTime)
	if err != nil {
		return invalidParamError(err, "start")
	}
	end, err := parseTimeParam(r, "end", maxTime)
	if err != nil {
		return invalidParamError(err, "end")
	}
	if end.Before(start) {
		err := errors.New("end timestamp must not be before start timestamp")
		return apiFuncResult{nil, &apiError{errorBadData, err}, nil, nil}
	}

	expr, err := parser.ParseExpr(r.FormValue("query"))
	if err != nil {
		return apiFuncResult{nil, &apiError{errorBadData, err}, nil, nil}
	}

	selectors := parser.ExtractSelectors(expr)
	if len(selectors) < 1 {
		return apiFuncResult{nil, nil, nil, nil}
	}

	ctx := r.Context()
	eq, err := api.exemplarQueryable.ExemplarQuerier(ctx)
	if err != nil {
		return apiFuncResult{nil, returnAPIError(err), nil, nil}
	}

	res, err := eq.Select(timestamp.FromTime(start), timestamp.FromTime(end), selectors...)
	if err != nil {
		return apiFuncResult{nil, returnAPIError(err), nil, nil}
	}

	return apiFuncResult{res, nil, nil, nil}
}

func (api *EngineAPI) metricMetadata(r *http.Request) apiFuncResult {
	metrics := map[string]map[apiMetadata]struct{}{}

	limit := -1
	if s := r.FormValue("limit"); s != "" {
		var err error
		if limit, err = strconv.Atoi(s); err != nil {
			return apiFuncResult{nil, &apiError{errorBadData, errors.New("limit must be a number")}, nil, nil}
		}
	}

	metric := r.FormValue("metric")
	for _, tt := range api.targetRetriever(r.Context()).TargetsActive() {
		for _, t := range tt {

			if metric == "" {
				for _, mm := range t.MetadataList() {
					m := apiMetadata{Type: mm.Type, Help: mm.Help, Unit: mm.Unit}
					ms, ok := metrics[mm.Metric]

					if !ok {
						ms = map[apiMetadata]struct{}{}
						metrics[mm.Metric] = ms
					}
					ms[m] = struct{}{}
				}
				continue
			}

			if md, ok := t.Metadata(metric); ok {
				m := apiMetadata{Type: md.Type, Help: md.Help, Unit: md.Unit}
				ms, ok := metrics[md.Metric]

				if !ok {
					ms = map[apiMetadata]struct{}{}
					metrics[md.Metric] = ms
				}
				ms[m] = struct{}{}
			}
		}
	}

	// Put the elements from the pseudo-set into a slice for marshaling.
	res := map[string][]apiMetadata{}
	for name, set := range metrics {
		if limit >= 0 && len(res) >= limit {
			break
		}

		s := []apiMetadata{}
		for metadata := range set {
			s = append(s, metadata)
		}
		res[name] = s
	}

	return apiFuncResult{res, nil, nil, nil}
}

func (api *EngineAPI) labelValues(r *http.Request) (result apiFuncResult) {
	ctx := r.Context()
	name := route.Param(ctx, "name")

	if !model.LabelNameRE.MatchString(name) {
		return apiFuncResult{nil, &apiError{errorBadData, errors.Errorf("invalid label name: %q", name)}, nil, nil}
	}

	start, err := parseTimeParam(r, "start", minTime)
	if err != nil {
		return invalidParamError(err, "start")
	}
	end, err := parseTimeParam(r, "end", maxTime)
	if err != nil {
		return invalidParamError(err, "end")
	}

	matcherSets, err := parseMatchersParam(r.Form["match[]"])
	if err != nil {
		return apiFuncResult{nil, &apiError{errorBadData, err}, nil, nil}
	}

	q, err := api.queryable.Querier(r.Context(), timestamp.FromTime(start), timestamp.FromTime(end))
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
		warnings storage.Warnings
	)
	if len(matcherSets) > 0 {
		var callWarnings storage.Warnings
		labelValuesSet := make(map[string]struct{})
		for _, matchers := range matcherSets {
			vals, callWarnings, err = q.LabelValues(name, matchers...)
			if err != nil {
				return apiFuncResult{nil, &apiError{errorExec, err}, warnings, closer}
			}
			warnings = append(warnings, callWarnings...)
			for _, val := range vals {
				labelValuesSet[val] = struct{}{}
			}
		}

		vals = make([]string, 0, len(labelValuesSet))
		for val := range labelValuesSet {
			vals = append(vals, val)
		}
	} else {
		vals, warnings, err = q.LabelValues(name)
		if err != nil {
			return apiFuncResult{nil, &apiError{errorExec, err}, warnings, closer}
		}

		if vals == nil {
			vals = []string{}
		}
	}

	slices.Sort(vals)

	return apiFuncResult{vals, nil, warnings, closer}
}

func (api *EngineAPI) labelNames(r *http.Request) apiFuncResult {
	start, err := parseTimeParam(r, "start", minTime)
	if err != nil {
		return invalidParamError(err, "start")
	}
	end, err := parseTimeParam(r, "end", maxTime)
	if err != nil {
		return invalidParamError(err, "end")
	}

	matcherSets, err := parseMatchersParam(r.Form["match[]"])
	if err != nil {
		return apiFuncResult{nil, &apiError{errorBadData, err}, nil, nil}
	}

	q, err := api.queryable.Querier(r.Context(), timestamp.FromTime(start), timestamp.FromTime(end))
	if err != nil {
		return apiFuncResult{nil, returnAPIError(err), nil, nil}
	}
	defer q.Close()

	var (
		names    []string
		warnings storage.Warnings
	)
	if len(matcherSets) > 0 {
		labelNamesSet := make(map[string]struct{})

		for _, matchers := range matcherSets {
			vals, callWarnings, err := q.LabelNames(matchers...)
			if err != nil {
				return apiFuncResult{nil, returnAPIError(err), warnings, nil}
			}

			warnings = append(warnings, callWarnings...)
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
		names, warnings, err = q.LabelNames()
		if err != nil {
			return apiFuncResult{nil, &apiError{errorExec, err}, warnings, nil}
		}
	}

	if names == nil {
		names = []string{}
	}
	return apiFuncResult{names, nil, warnings, nil}
}

func (api *EngineAPI) query(r *http.Request) (result apiFuncResult) {
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
		ctx, cancel = context.WithTimeout(ctx, timeout)
		defer cancel()
	}

	opts, err := extractQueryOpts(r)
	if err != nil {
		return apiFuncResult{nil, &apiError{errorBadData, err}, nil, nil}
	}
	qry, err := api.queryEngine.NewInstantQuery(ctx, api.queryable, opts, r.FormValue("query"), ts)
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
		sr = defaultStatsRenderer
	}
	qs := sr(ctx, qry.Stats(), r.FormValue("stats"))

	return apiFuncResult{&queryData{
		ResultType: res.Value.Type(),
		Result:     res.Value,
		Stats:      qs,
	}, nil, res.Warnings, qry.Close}
}

func (api *EngineAPI) series(r *http.Request) (result apiFuncResult) {
	if err := r.ParseForm(); err != nil {
		return apiFuncResult{nil, &apiError{errorBadData, errors.Wrapf(err, "error parsing form values")}, nil, nil}
	}
	if len(r.Form["match[]"]) == 0 {
		return apiFuncResult{nil, &apiError{errorBadData, errors.New("no match[] parameter provided")}, nil, nil}
	}

	start, err := parseTimeParam(r, "start", minTime)
	if err != nil {
		return invalidParamError(err, "start")
	}
	end, err := parseTimeParam(r, "end", maxTime)
	if err != nil {
		return invalidParamError(err, "end")
	}

	matcherSets, err := parseMatchersParam(r.Form["match[]"])
	if err != nil {
		return invalidParamError(err, "match[]")
	}

	q, err := api.queryable.Querier(r.Context(), timestamp.FromTime(start), timestamp.FromTime(end))
	if err != nil {
		return apiFuncResult{nil, returnAPIError(err), nil, nil}
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

	hints := &storage.SelectHints{
		Start: timestamp.FromTime(start),
		End:   timestamp.FromTime(end),
		Func:  "series", // There is no series function, this token is used for lookups that don't need samples.
	}
	var set storage.SeriesSet

	if len(matcherSets) > 1 {
		var sets []storage.SeriesSet
		for _, mset := range matcherSets {
			// We need to sort this select results to merge (deduplicate) the series sets later.
			s := q.Select(true, hints, mset...)
			sets = append(sets, s)
		}
		set = storage.NewMergeSeriesSet(sets, storage.ChainedSeriesMerge)
	} else {
		// At this point at least one match exists.
		set = q.Select(false, hints, matcherSets[0]...)
	}

	metrics := []labels.Labels{}
	for set.Next() {
		metrics = append(metrics, set.At().Labels())
	}

	warnings := set.Warnings()
	if set.Err() != nil {
		return apiFuncResult{nil, returnAPIError(set.Err()), warnings, closer}
	}

	return apiFuncResult{metrics, nil, warnings, closer}
}

func (api *EngineAPI) queryRange(r *http.Request) (result apiFuncResult) {
	start, err := parseTime(r.FormValue("start"))
	if err != nil {
		return invalidParamError(err, "start")
	}
	end, err := parseTime(r.FormValue("end"))
	if err != nil {
		return invalidParamError(err, "end")
	}
	if end.Before(start) {
		return invalidParamError(errors.New("end timestamp must not be before start time"), "end")
	}

	step, err := parseDuration(r.FormValue("step"))
	if err != nil {
		return invalidParamError(err, "step")
	}

	if step <= 0 {
		return invalidParamError(errors.New("zero or negative query resolution step widths are not accepted. Try a positive integer"), "step")
	}

	// For safety, limit the number of returned points per timeseries.
	// This is sufficient for 60s resolution for a week or 1h resolution for a year.
	if end.Sub(start)/step > 11000 {
		err := errors.New("exceeded maximum resolution of 11,000 points per timeseries. Try decreasing the query resolution (?step=XX)")
		return apiFuncResult{nil, &apiError{errorBadData, err}, nil, nil}
	}

	ctx := r.Context()
	if to := r.FormValue("timeout"); to != "" {
		var cancel context.CancelFunc
		timeout, err := parseDuration(to)
		if err != nil {
			return invalidParamError(err, "timeout")
		}

		ctx, cancel = context.WithTimeout(ctx, timeout)
		defer cancel()
	}

	opts, err := extractQueryOpts(r)
	if err != nil {
		return apiFuncResult{nil, &apiError{errorBadData, err}, nil, nil}
	}
	qry, err := api.queryEngine.NewRangeQuery(ctx, api.queryable, opts, r.FormValue("query"), start, end, step)
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
		sr = defaultStatsRenderer
	}
	qs := sr(ctx, qry.Stats(), r.FormValue("stats"))

	return apiFuncResult{&queryData{
		ResultType: res.Value.Type(),
		Result:     res.Value,
		Stats:      qs,
	}, nil, res.Warnings, qry.Close}
}

func (api *EngineAPI) respond(w http.ResponseWriter, data interface{}, warnings storage.Warnings) {
	statusMessage := statusSuccess
	var warningStrings []string
	for _, warning := range warnings {
		warningStrings = append(warningStrings, warning.Error())
	}
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	b, err := json.Marshal(&response{
		Status:   statusMessage,
		Data:     data,
		Warnings: warningStrings,
	})
	if err != nil {
		_ = level.Error(api.logger).Log("msg", "error marshaling json response", "err", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if n, err := w.Write(b); err != nil {
		_ = level.Error(api.logger).Log("msg", "error writing response", "bytesWritten", n, "err", err)
	}
}

func (api *EngineAPI) respondError(w http.ResponseWriter, apiErr *apiError, data interface{}) {
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	b, err := json.Marshal(&response{
		Status:    statusError,
		ErrorType: apiErr.typ,
		Error:     apiErr.err.Error(),
		Data:      data,
	})
	if err != nil {
		_ = level.Error(api.logger).Log("msg", "error marshaling json response", "err", err)
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
	default:
		code = http.StatusInternalServerError
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if n, err := w.Write(b); err != nil {
		_ = level.Error(api.logger).Log("msg", "error writing response", "bytesWritten", n, "err", err)
	}
}
