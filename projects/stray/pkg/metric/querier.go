package metric

import (
	"context"
	"time"

	kitlog "github.com/go-kit/log"
	"github.com/go-kit/log/level"

	"github.com/prometheus/common/model"
	"github.com/prometheus/prometheus/config"
	"github.com/prometheus/prometheus/promql"
	"github.com/prometheus/prometheus/storage"
	"go.uber.org/atomic"
)

type (
	PromQuerier struct {
		queryEngine *promql.Engine
		log         kitlog.Logger
	}

	PromQuerierOpts struct {
		Log           kitlog.Logger
		MaxMaxSamples int
		Timeout       time.Duration
	}

	safePromQLNoStepSubqueryInterval struct {
		value atomic.Int64
	}
)

func durationToInt64Millis(d time.Duration) int64 {
	return int64(d / time.Millisecond)
}

func (i *safePromQLNoStepSubqueryInterval) Set(ev model.Duration) {
	i.value.Store(durationToInt64Millis(time.Duration(ev)))
}

func (i *safePromQLNoStepSubqueryInterval) Get(int64) int64 {
	return i.value.Load()
}

func makePromQLEngOpts(opts *PromQuerierOpts, logger kitlog.Logger) promql.EngineOpts {
	noStepSubqueryInterval := &safePromQLNoStepSubqueryInterval{}
	noStepSubqueryInterval.Set(config.DefaultGlobalConfig.EvaluationInterval)
	ret := promql.EngineOpts{
		Logger:                   logger,
		NoStepSubqueryIntervalFn: noStepSubqueryInterval.Get,
		EnableAtModifier:         true,
		EnableNegativeOffset:     true,
		EnablePerStepStats:       true,
		LookbackDelta:            time.Duration(5 * time.Minute),
		Reg:                      nil,
		MaxSamples:               opts.MaxMaxSamples,
		Timeout:                  opts.Timeout,
	}
	if ret.Timeout < QL_TIMEOUT {
		ret.Timeout = QL_TIMEOUT
	}
	if ret.MaxSamples < QL_MAX_SAMPLES {
		ret.MaxSamples = QL_MAX_SAMPLES
	}
	return ret
}

func NewQuerier(opts PromQuerierOpts) (*PromQuerier, error) {
	logger := kitlog.With(opts.Log, LOG_COMPONENT_KEY, COMPONENT_PMQL)
	return &PromQuerier{
		queryEngine: promql.NewEngine(makePromQLEngOpts(&opts, logger)),
		log:         logger,
	}, nil
}

func (pmq *PromQuerier) GetPmQL() *promql.Engine {
	return pmq.queryEngine
}

func (pmq *PromQuerier) SetQueryLogger(l promql.QueryLogger) {
	_ = level.Debug(pmq.log).Log("msg", "setQueryLogger")
	pmq.queryEngine.SetQueryLogger(l)
}

func (pmq *PromQuerier) NewInstantQuery(ctx context.Context, q storage.Queryable,
	opts *promql.QueryOpts, qs string, ts time.Time) (promql.Query, error) {
	_ = level.Debug(pmq.log).Log("msg", "inst query", "qs", qs, "ts", ts)
	return pmq.queryEngine.NewInstantQuery(ctx, q, opts, qs, ts)
}

func (pmq *PromQuerier) NewRangeQuery(ctx context.Context, q storage.Queryable,
	opts *promql.QueryOpts, qs string, start, end time.Time,
	interval time.Duration) (promql.Query, error) {
	_ = level.Debug(pmq.log).Log("msg", "rang query",
		"qs", qs, "start", start, "end", end, "interval", interval)
	return pmq.queryEngine.NewRangeQuery(ctx, q, opts, qs, start, end, interval)
}
