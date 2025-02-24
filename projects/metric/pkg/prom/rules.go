package prom

import (
	"context"
	"time"

	kitlog "github.com/go-kit/log"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/prometheus/rules"
)

type (
	PromRulesOpts struct {
		LoggerHandler LogAdapterHandler

		// Logger       kitlog.Logger
		MetricEngine *Engine
	}

	PromRules struct {
		logger kitlog.Logger

		rulesCtx context.Context
		rulesMgr *rules.Manager
	}
)

func onNotify(ctx context.Context, expr string, alerts ...*rules.Alert) {
	// TODO send alert to notification ?
	// rules.SendAlerts(notifierManager, cfg.web.ExternalURL.String()),
}

func NewPromRules(opts PromRulesOpts) (*PromRules, error) {
	ctx := context.TODO()
	promLog := NewLogAdapter(opts.LoggerHandler)
	promStorage := opts.MetricEngine.GetStorage()
	promQuerier := opts.MetricEngine.GetQuerier()
	ruleManager := rules.NewManager(&rules.ManagerOptions{
		Appendable:      promStorage,
		Queryable:       promStorage,
		QueryFunc:       rules.EngineQueryFunc(promQuerier.GetPmQL(), promStorage),
		Context:         ctx,
		NotifyFunc:      onNotify,
		Registerer:      prometheus.DefaultRegisterer,
		Logger:          kitlog.With(promLog, "component", "rules"),
		OutageTolerance: time.Duration(time.Hour * 1),
		ForGracePeriod:  time.Duration(time.Minute * 10),
		ResendDelay:     time.Duration(time.Minute * 1),
		// TODO Set Rules group loader at here
		// GroupLoader: ,
	})
	return &PromRules{
		logger:   promLog,
		rulesCtx: ctx,
		rulesMgr: ruleManager,
	}, nil
}

func (r *PromRules) Run() {
	r.rulesMgr.Run()
}
