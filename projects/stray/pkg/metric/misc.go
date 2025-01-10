package metric

import (
	"log/slog"

	"github.com/grafana/regexp"
	"github.com/prometheus/common/model"
	"github.com/prometheus/prometheus/discovery"
	"github.com/prometheus/prometheus/discovery/targetgroup"
	"github.com/prometheus/prometheus/model/relabel"
	"github.com/samber/lo"
)

type (
	StaticTargetGroup struct {
		Source    string
		Addresses []string
	}
)

func makeComponentLog(log *slog.Logger, component string) *slog.Logger {
	return lo.If(log == nil, slog.Default()).Else(log).With(LOG_COMPONENT_KEY,
		lo.If(component == "", "default").Else(component))
}

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

func compileCORSRegexString(s string) (*regexp.Regexp, error) {
	r, err := relabel.NewRegexp(s)
	if err != nil {
		slog.Error("compile cors reg", slog.Any("err", err), slog.String("src", s))
		return nil, err
	}
	return r.Regexp, nil
}
