package util

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/common/model"
	"github.com/prometheus/prometheus/discovery"
	"github.com/prometheus/prometheus/discovery/targetgroup"
)

type (
	PromDiscovery struct {
		ctx     context.Context
		promLog *slog.Logger

		mgr *discovery.Manager
	}
)

func staticConfig(addrs ...string) discovery.StaticConfig {
	var cfg discovery.StaticConfig
	for i, addr := range addrs {
		cfg = append(cfg, &targetgroup.Group{
			Source: fmt.Sprint(i),
			Targets: []model.LabelSet{
				{model.AddressLabel: model.LabelValue(addr)},
			},
		})
	}
	return cfg
}

func NewPromDiscovery(ctx context.Context) (*PromDiscovery, error) {
	promLog := slog.Default().With("component", "discovery")
	sdMetrics, err := discovery.CreateAndRegisterSDMetrics(prometheus.DefaultRegisterer)
	if err != nil {
		slog.Error("failed to register service discovery metrics", "err", err)
		return nil, err
	}
	mgr := discovery.NewManager(ctx, promLog, prometheus.DefaultRegisterer, sdMetrics)
	return &PromDiscovery{ctx: ctx, mgr: mgr, promLog: promLog}, nil
}

func (d *PromDiscovery) Run() {
	go func() {
		d.Run()
	}()
}

func (d *PromDiscovery) ApplyConfig() {
	// XXX TODO Load the addresses from parameters
	discoveryCfg := map[string]discovery.Configs{
		"job1": {staticConfig("172.24.6.50:9100")},
	}
	d.mgr.ApplyConfig(discoveryCfg)
}
