package metric

import (
	"context"
	"log/slog"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/prometheus/discovery"
	"github.com/samber/lo"
)

type (
	PromDiscoveryOpts struct {
		PromLog *slog.Logger
		PromReg *prometheus.Registerer
	}

	PromDiscovery struct {
		ctx context.Context
		mgr *discovery.Manager
	}

	StaticDiscoveryConfig struct {
		JobName string
		Targets []StaticTargetGroup
	}
)

func NewPromDiscovery(ctx context.Context, opts PromDiscoveryOpts) (*PromDiscovery, error) {
	compLog := makeComponentLog(opts.PromLog, COMPONENT_DISCOVERY)
	sdMetrics, err := discovery.CreateAndRegisterSDMetrics(
		lo.If(opts.PromReg == nil, prometheus.DefaultRegisterer).
			Else(*opts.PromReg))
	if err != nil {
		slog.Error("failed to register service discovery metrics", "err", err)
		return nil, err
	}
	mgr := discovery.NewManager(ctx, compLog, prometheus.DefaultRegisterer, sdMetrics)
	return &PromDiscovery{ctx: ctx, mgr: mgr}, nil
}

func (d *PromDiscovery) Get() *discovery.Manager {
	return d.mgr
}

func (d *PromDiscovery) Run() error {
	slog.Debug("launching SD component")
	return d.mgr.Run()
}

func (d *PromDiscovery) ApplyStaticTargetGroup(groups ...StaticDiscoveryConfig) {
	discoveryCfg := make(map[string]discovery.Configs)
	lo.ForEach(groups, func(grp StaticDiscoveryConfig, _ int) {
		discoveryCfg[grp.JobName] = []discovery.Config{
			makeDiscoveryStaticConfig(grp.Targets...),
		}
	})
	slog.Debug("sd apply static config",
		slog.Any("grp", groups),
		slog.Any("cfg", discoveryCfg))
	d.mgr.ApplyConfig(discoveryCfg)
}
