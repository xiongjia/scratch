package metric

import (
	"log/slog"
	"time"

	"github.com/prometheus/common/model"
	"github.com/prometheus/prometheus/config"
	"github.com/prometheus/prometheus/promql"
	"go.uber.org/atomic"
)

type (
	PromQuerier struct {
		QueryEngine *promql.Engine
	}

	PromQuerierOpts struct {
		Log           *slog.Logger
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

func makePromQLEngOpts(opts *PromQuerierOpts) promql.EngineOpts {
	noStepSubqueryInterval := &safePromQLNoStepSubqueryInterval{}
	noStepSubqueryInterval.Set(config.DefaultGlobalConfig.EvaluationInterval)
	ret := promql.EngineOpts{
		Logger:                   makeComponentLog(opts.Log, COMPONENT_PMQL),
		NoStepSubqueryIntervalFn: noStepSubqueryInterval.Get,
		EnableAtModifier:         true,
		EnableNegativeOffset:     true,
		EnablePerStepStats:       true,
		LookbackDelta:            time.Duration(5 * time.Minute),
		EnableDelayedNameRemoval: true,
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
	return &PromQuerier{QueryEngine: promql.NewEngine(makePromQLEngOpts(&opts))}, nil
}
