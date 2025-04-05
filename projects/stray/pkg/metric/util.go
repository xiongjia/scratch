package metric

import (
	"sync"
	"time"

	"github.com/prometheus/common/model"
	"github.com/prometheus/prometheus/discovery"
	"github.com/prometheus/prometheus/discovery/targetgroup"
	"github.com/prometheus/prometheus/model/histogram"
	"github.com/prometheus/prometheus/storage"
	"github.com/prometheus/prometheus/tsdb/chunkenc"

	"github.com/samber/lo"
)

type (
	WaitGroup struct {
		wg sync.WaitGroup
	}

	StaticTargetGroup struct {
		Source    string
		Addresses []string
	}

	FloatSample struct {
		t int64
		f float64
	}

	PromChunkSeriesSet struct {
		idx    int
		series []storage.ChunkSeries
	}
)

func NewWaitGroup() *WaitGroup {
	return &WaitGroup{}
}

func (h *WaitGroup) Go(f func()) {
	h.wg.Add(1)
	go func() {
		defer h.wg.Done()
		f()
	}()
}

func (h *WaitGroup) Wait() {
	h.wg.Wait()
}

func MaxDuration(d1, d2 time.Duration) time.Duration {
	if d1 > d2 {
		return d1
	} else {
		return d2
	}
}

func MaxNum(n1, n2 int) int {
	if n1 > n2 {
		return n1
	} else {
		return n2
	}
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

func findAddressesInStaticTargetGroup(grp []StaticTargetGroup, src string) []string {
	for _, itr := range grp {
		if itr.Source == src {
			return itr.Addresses
		}
	}
	return []string{}
}

func findStaticTargetConf(c []StaticDiscoveryConfig, jobName string) []StaticTargetGroup {
	for _, itr := range c {
		if itr.JobName == jobName {
			return itr.Targets
		}
	}
	return []StaticTargetGroup{}
}

func compareStaticTarget(t1, t2 []string) bool {
	if len(t1) != len(t2) {
		return false
	}
	for _, itr1 := range t1 {
		found := false
		for _, itr2 := range t2 {
			if itr1 == itr2 {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

func compareStaticTargetGroup(g1, g2 []StaticTargetGroup) bool {
	if len(g1) != len(g2) {
		return false
	}
	for _, itr := range g1 {
		g2Addr := findAddressesInStaticTargetGroup(g2, itr.Source)
		if !compareStaticTarget(itr.Addresses, g2Addr) {
			return false
		}
	}
	return true
}

func compareStaticTargetConf(c1, c2 []StaticDiscoveryConfig) bool {
	if len(c1) != len(c2) {
		return false
	}

	for _, itr := range c1 {
		c2Targets := findStaticTargetConf(c2, itr.JobName)
		if !compareStaticTargetGroup(itr.Targets, c2Targets) {
			return false
		}
	}
	return true
}

func findScrapeJobs(src []ScrapeJob, jobName string) *ScrapeJob {
	for _, itr := range src {
		if itr.JobName == jobName {
			return &itr
		}
	}
	return nil
}

func compareScrapeJob(j1, j2 *ScrapeJob) bool {
	if j1 == nil || j2 == nil {
		return false
	}
	if j1.JobName == j2.JobName &&
		j1.MetricsPath == j2.MetricsPath &&
		j1.Scheme == j2.Scheme {
		return true
	} else {
		return false
	}
}

func compareScrapeJobs(j1, j2 []ScrapeJob) bool {
	if len(j1) != len(j2) {
		return false
	}
	for _, itr1 := range j1 {
		itr2 := findScrapeJobs(j2, itr1.JobName)
		if itr2 == nil {
			return false
		}
		if !compareScrapeJob(&itr1, itr2) {
			return false
		}
	}
	return true
}

func (s FloatSample) T() int64 {
	return s.t
}

func (s FloatSample) F() float64 {
	return s.f
}

func (s FloatSample) H() *histogram.Histogram {
	panic("H() called for fSample")
}

func (s FloatSample) FH() *histogram.FloatHistogram {
	panic("FH() called for fSample")
}

func (s FloatSample) Type() chunkenc.ValueType {
	return chunkenc.ValFloat
}

func NewPromChunkSeriesSet(series ...storage.ChunkSeries) storage.ChunkSeriesSet {
	return &PromChunkSeriesSet{idx: -1, series: series}
}

func (m *PromChunkSeriesSet) Next() bool {
	m.idx++
	return m.idx < len(m.series)
}

func (m *PromChunkSeriesSet) At() storage.ChunkSeries {
	return m.series[m.idx]
}

func (m *PromChunkSeriesSet) Err() error {
	return nil
}

func (m *PromChunkSeriesSet) Warnings() storage.Warnings {
	return nil
}
