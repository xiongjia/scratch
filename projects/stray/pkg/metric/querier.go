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

func NewQuerier() (*PromQuerier, error) {
	noStepSubqueryInterval := &safePromQLNoStepSubqueryInterval{}
	noStepSubqueryInterval.Set(config.DefaultGlobalConfig.EvaluationInterval)

	engLog := slog.Default().With("component", "pmql")
	queryEngine := promql.NewEngine(promql.EngineOpts{
		Logger:                   engLog,
		Reg:                      nil,
		MaxSamples:               5000000,
		Timeout:                  100 * time.Second,
		NoStepSubqueryIntervalFn: noStepSubqueryInterval.Get,
		EnableAtModifier:         true,
		EnableNegativeOffset:     true,
		EnablePerStepStats:       true,
		LookbackDelta:            time.Duration(5 * time.Minute),
		EnableDelayedNameRemoval: true,
	})
	return &PromQuerier{QueryEngine: queryEngine}, nil
}
