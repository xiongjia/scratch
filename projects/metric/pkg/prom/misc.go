package prom

import (
	"github.com/prometheus/common/model"
	"github.com/prometheus/prometheus/discovery"
	"github.com/prometheus/prometheus/discovery/targetgroup"

	"github.com/samber/lo"
)

type (
	StaticTargetGroup struct {
		Source    string
		Addresses []string
	}
)

func makeDiscoveryStaticConfig(groups ...StaticTargetGroup) discovery.StaticConfig {
	var cfg discovery.StaticConfig
	lo.ForEach(groups, func(grp StaticTargetGroup, _ int) {
		lo.ForEach(grp.Addresses, func(addr string, _ int) {
			cfg = append(cfg, &targetgroup.Group{
				Source: grp.Source,
				Targets: []model.LabelSet{
					{model.AddressLabel: model.LabelValue(addr)},
				},
			})
		})
	})
	return cfg
}
