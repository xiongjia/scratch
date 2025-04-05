package metric

import (
	"context"

	kitlog "github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/prometheus/prometheus/discovery"
	"github.com/samber/lo"
)

type (
	PromDiscoveryOpts struct {
		Logger kitlog.Logger
	}

	PromDiscovery struct {
		logger kitlog.Logger

		ctx context.Context
		mgr *discovery.Manager
	}

	StaticDiscoveryConfig struct {
		JobName string
		Targets []StaticTargetGroup
	}
)

func NewPromDiscovery(ctx context.Context, opts PromDiscoveryOpts) (*PromDiscovery, error) {
	_ = level.Debug(opts.Logger).Log("msg", "create prom discovery", "opts", opts)
	mgr := discovery.NewManager(ctx, kitlog.With(opts.Logger, LOG_COMPONENT_KEY, COMPONENT_DISCOVERY))
	return &PromDiscovery{logger: opts.Logger, ctx: ctx, mgr: mgr}, nil
}

func (d *PromDiscovery) Get() *discovery.Manager {
	return d.mgr
}

func (d *PromDiscovery) Run() error {
	_ = level.Debug(d.logger).Log("msg", "launching SD component")
	return d.mgr.Run()
}

func (d *PromDiscovery) ApplyStaticTargetGroup(groups []StaticDiscoveryConfig) error {
	discoveryCfg := make(map[string]discovery.Configs)
	lo.ForEach(groups, func(grp StaticDiscoveryConfig, _ int) {
		discoveryCfg[grp.JobName] = []discovery.Config{
			makeDiscoveryStaticConfig(grp.Targets...),
		}
	})
	_ = level.Debug(d.logger).Log("msg", "sd apply static config", "grp", groups, "cfg", discoveryCfg)
	err := d.mgr.ApplyConfig(discoveryCfg)
	if err != nil {
		_ = level.Error(d.logger).Log("msg", "service discovery apply config", "err", err)
	}
	return err
}
