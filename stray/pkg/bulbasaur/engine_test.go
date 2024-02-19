package bulbasaur_test

import (
	"context"
	"fmt"
	"html/template"
	"log/slog"
	"os"
	"stray/pkg/bulbasaur"
	"stray/pkg/dugtrio"
	"testing"
	"time"

	"github.com/oklog/run"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/prometheus/config"
	"github.com/prometheus/prometheus/discovery"
	"github.com/prometheus/prometheus/discovery/targetgroup"
	"github.com/prometheus/prometheus/scrape"
	"github.com/prometheus/prometheus/tsdb"
	"gopkg.in/yaml.v2"
)

func TestEngine(t *testing.T) {
	logger := dugtrio.NewSlogForTesting(t, dugtrio.SLogBaseOptions{Level: slog.LevelDebug})

	engine, err := bulbasaur.NewEngine(bulbasaur.EngineOptions{Logger: logger})
	if err != nil {
		return
	}
	engine.Test()
}

func loadCfg(c string) *config.Config {
	cfg := &config.Config{}
	if err := yaml.UnmarshalStrict([]byte(c), cfg); err != nil {
		fmt.Printf("err: %s\n", err)
	}
	return cfg
}

func TestEng2(t *testing.T) {
	// creating scrape
	testRegistry := prometheus.NewRegistry()
	opts := &scrape.Options{}
	mgr, err := scrape.NewManager(opts, nil, nil, testRegistry)
	if err != nil {
		return
	}

	cfg := loadCfg(`
    scrape_configs:
    - job_name: job1
      static_configs:
      - targets: ["foo:9090"]
    - job_name: job2
      static_configs:
      - targets: ["foo:9091", "foo:9092"]
    `)
	fmt.Printf("cfg: %v\n", cfg)

	mgr.ApplyConfig(cfg)
	p := mgr.ScrapePools()
	fmt.Printf("pools: %v\n", p)

	ts := make(chan map[string][]*targetgroup.Group)
	fmt.Printf("ts: %v\n", ts)
	go mgr.Run(ts)

	time.Sleep(60 * time.Second)
	p = mgr.ScrapePools()
	fmt.Printf("pools: %v\n", p)
}

func TestRunGrp(t *testing.T) {
	var runGrp run.Group
	runGrp.Add(func() error {
		time.Sleep(3 * time.Second)
		fmt.Printf("func1\n")
		// time.Sleep(10 * time.Second)
		// fmt.Printf("func1 ended\n")
		return nil
	}, func(err error) {
		fmt.Printf("func1 stopped\n")
		if err != nil {
			fmt.Printf("func1 Error : %s\n", err)
		}
	})
	runGrp.Add(func() error {
		time.Sleep(10 * time.Second)
		fmt.Printf("func2\n")
		return fmt.Errorf("func2 error string")
		// return nil
	}, func(err error) {
		fmt.Printf("func2 stopped\n")
		if err != nil {
			fmt.Printf("func2 Error : %s\n", err)
		}
	})

	err := runGrp.Run()
	if err != nil {
		fmt.Printf("Error : %s\n", err)
	}

	fmt.Printf("test passed\n")
}

func TestSd(t *testing.T) {
	ctxScrape, _ := context.WithCancel(context.Background())
	discMgr := discovery.NewManager(ctxScrape, nil)

	cfg := loadCfg(`
    scrape_configs:
    - job_name: jobMain
      static_configs:
      - targets: ["172.24.10.193:9100"]
    `)

	c := make(map[string]discovery.Configs)
	scfgs, err := cfg.GetScrapeConfigs()
	if err != nil {
		return
	}
	for _, v := range scfgs {
		c[v.JobName] = v.ServiceDiscoveryConfigs
	}
	discMgr.ApplyConfig(c)

	db, err := tsdb.Open("c:/wrk/tmp/tsdb1", nil, nil, tsdb.DefaultOptions(), nil)
	if err != nil {
		return
	}
	testRegistry := prometheus.NewRegistry()
	scrapeOpts := &scrape.Options{}
	scrapMgr, err := scrape.NewManager(scrapeOpts, nil, db, testRegistry)
	if err != nil {
		return
	}
	scrapMgr.ApplyConfig(cfg)

	var runGrp run.Group

	runGrp.Add(func() error {
		fmt.Printf("discover mgr started\n")
		err = discMgr.Run()
		if err != nil {
			fmt.Printf("error %v\n", err)
		}
		return err
	}, func(err error) {
		fmt.Printf("discover mgr stopped\n")
	})

	runGrp.Add(func() error {
		fmt.Printf("scrap mgr started\n")
		err := scrapMgr.Run(discMgr.SyncCh())
		if err != nil {
			fmt.Printf("error %v\n", err)
		}
		return nil
	}, func(err error) {
		fmt.Printf("scrapMgr stopped\n")
	})

	runGrp.Run()
	fmt.Printf("test passed")
}

func TestTemplate(t *testing.T) {
	const tpl = `
<!DOCTYPE html>
<html>
	<head>
		<meta charset="UTF-8">
		<title>{{.Title}}</title>
	</head>
	<body>
		{{range .Items}}<div>{{ . }}</div>{{else}}<div><strong>no rows</strong></div>{{end}}
	</body>
</html>`

	tmpl, err := template.New("webpage").Parse(tpl)
	if err != nil {
		return
	}

	data := struct {
		Title string
		Items []string
	}{
		Title: "My page",
		Items: []string{},
	}
	tmpl.Execute(os.Stdout, data)
}
