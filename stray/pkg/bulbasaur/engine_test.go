package bulbasaur_test

import (
	"fmt"
	"log/slog"
	"stray/pkg/bulbasaur"
	"stray/pkg/dugtrio"
	"testing"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/prometheus/config"
	"github.com/prometheus/prometheus/discovery/targetgroup"
	"github.com/prometheus/prometheus/scrape"
)

func TestEngine(t *testing.T) {
	logger := dugtrio.NewSlogForTesting(t, dugtrio.SLogBaseOptions{Level: slog.LevelDebug})

	engine, err := bulbasaur.NewEngine(bulbasaur.EngineOptions{Logger: logger})
	if err != nil {
		return
	}
	engine.Test()
}

func TestEng2(t *testing.T) {
	cfgText1 := `
    scrape_configs:
     - job_name: job1
       static_configs:
       - targets: ["127.0.0.1:2112"]
    `
	cfg2, err := config.Load(cfgText1, false, nil)
	if err != nil {
		t.Fatalf("Unable to load YAML config: %s", err)
	}
	fmt.Printf("%v\n", cfg2)

	testRegistry := prometheus.NewRegistry()
	opts := &scrape.Options{}
	scrapeMgr, err := scrape.NewManager(opts, nil, nil, testRegistry)
	if err != nil {
		return
	}

	err = scrapeMgr.ApplyConfig(cfg2)
	if err != nil {
		fmt.Printf("err %s", err.Error())
	}

	ts := make(chan map[string][]*targetgroup.Group)
	go func() {
		err := scrapeMgr.Run(ts)
		if err != nil {
			fmt.Printf("mgr error %s\n", err.Error())
		}
		fmt.Printf("stopped\n")
	}()

	defer scrapeMgr.Stop()

	tgSent := make(map[string][]*targetgroup.Group)
	tgSent["job1"] = []*targetgroup.Group{
		{
			Source: "job1",
		},
	}
	ts <- tgSent

	pools := scrapeMgr.ScrapePools()
	fmt.Printf("Pools %v\n", pools)

	time.Sleep(60 * time.Second)

	pools = scrapeMgr.ScrapePools()
	fmt.Printf("Pools %v\n", pools)

	time.Sleep(240 * time.Second)
}
