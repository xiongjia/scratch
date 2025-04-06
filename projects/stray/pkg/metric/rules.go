package metric

import (
	"context"
	"net/http"
	"time"

	kitlog "github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/common/model"
	"github.com/prometheus/prometheus/model/labels"
	"github.com/prometheus/prometheus/model/rulefmt"
	"github.com/prometheus/prometheus/notifier"
	"github.com/prometheus/prometheus/promql/parser"
	"github.com/prometheus/prometheus/rules"
	"gopkg.in/yaml.v3"
)

type (
	PromRulesOpts struct {
		logger kitlog.Logger

		PromStorage *PromStorage
		PromQuerier *PromQuerier
	}

	PromRules struct {
		logger kitlog.Logger

		rulesCtx context.Context
		rulesMgr *rules.Manager

		notifierMgr *notifier.Manager
	}

	MetricGroupLoader struct {
		logger kitlog.Logger
	}

	MetricAlertSender struct {
		logger kitlog.Logger
	}
)

func makeRulesGroups() []rulefmt.RuleGroup {
	result := make([]rulefmt.RuleGroup, 0)
	rule1 := rulefmt.RuleNode{}
	// expr: r_process_open_fds{instance="127.0.0.1:8897"} > 2000
	err := yaml.Unmarshal([]byte(`
alert: InstanceDown
expr: up{instance="127.0.0.1:8897"} == 1
for: 1m
labels:
    severity: critical
annotations:
    description: Host {{ $labels.instance }} Down
    summary: Host {{ $labels.instance }} of {{ $labels.job }} is  Down
`), &rule1)
	if err != nil {
		return result
	}

	result = append(result, rulefmt.RuleGroup{
		Name:     "rule1",
		Interval: model.Duration(20 * time.Second),
		Limit:    0,
		Rules:    []rulefmt.RuleNode{rule1},
	})
	return result
}

func (l *MetricGroupLoader) Load(identifier string) (*rulefmt.RuleGroups, []error) {
	_ = level.Debug(l.logger).Log("msg", "load rule group", "id", identifier)
	return &rulefmt.RuleGroups{Groups: makeRulesGroups()}, nil
}

func (l *MetricGroupLoader) Parse(query string) (parser.Expr, error) {
	_ = level.Debug(l.logger).Log("msg", "parse rule", "query", query)
	return parser.ParseExpr(query)
}

func NewPromRules(opts PromRulesOpts) (*PromRules, error) {
	ctx := context.Background()

	notifierManager := notifier.NewManager(&notifier.Options{
		QueueCapacity: 0,
		Do: func(ctx context.Context, client *http.Client, req *http.Request) (*http.Response, error) {
			_ = level.Debug(opts.logger).Log("msg", "notif")
			return &http.Response{StatusCode: 200}, nil
		},
	}, kitlog.With(opts.logger, "component", "notifier"))

	alertSender := &MetricAlertSender{logger: opts.logger}
	result := &PromRules{logger: opts.logger, rulesCtx: ctx, notifierMgr: notifierManager}
	result.rulesMgr = rules.NewManager(&rules.ManagerOptions{
		Appendable: opts.PromStorage,
		Queryable:  opts.PromStorage,
		QueryFunc:  rules.EngineQueryFunc(opts.PromQuerier.queryEngine, opts.PromStorage),
		Context:    ctx,
		NotifyFunc: rules.SendAlerts(alertSender, "test"),
		// func(ctx context.Context, expr string, alerts ...*rules.Alert) {
		// 	result.onNotifyEvent(ctx, expr, alerts...)
		// },

		Registerer:      prometheus.DefaultRegisterer,
		Logger:          kitlog.With(opts.logger, "component", "rules"),
		OutageTolerance: time.Duration(time.Hour * 1),
		ForGracePeriod:  time.Duration(time.Minute * 10),
		ResendDelay:     time.Duration(time.Minute * 1),
		GroupLoader:     &MetricGroupLoader{logger: opts.logger},
	})
	return result, nil
}

func (s *MetricAlertSender) Send(alerts ...*notifier.Alert) {
	_ = level.Debug(s.logger).Log("msg", "on notify", "alerts", alerts)
}

func (r *PromRules) LoadRuleGroups() {
	// r.rulesMgr.LoadGroups(60*time.Second, labels.EmptyLabels(), "", nil, []string{"general"}...)
	r.rulesMgr.Update(60*time.Second, []string{"general"}, labels.EmptyLabels(), "", nil)

	alerts := r.rulesMgr.AlertingRules()
	_ = level.Debug(r.logger).Log("msg", "dbg", "alerts", alerts)

	grp := r.rulesMgr.RuleGroups()
	_ = level.Debug(r.logger).Log("msg", "dbg", "grp", grp)
}

func (r *PromRules) onNotifyEvent(ctx context.Context, expr string, alerts ...*rules.Alert) {
	_ = level.Debug(r.logger).Log("msg", "on notify", "expr", expr)
	if len(alerts) == 0 {
		return
	}
	_ = level.Debug(r.logger).Log("msg", "on notify", "expr", expr, "alert", alerts)

	// rules.SendAlerts(r.rulesMgr., "test")
	// r.rulesMgr.SendAlerts(notifierManager, cfg.web.ExternalURL.String()),
}

func (r *PromRules) Run() {
	r.rulesMgr.Run()
}
