package main

import (
	"flag"
	"log/slog"
	"metric/pkg/buildmgr"
)

var (
	projects = []*buildmgr.BinaryBuildOption{
		{
			Output:       "metric",
			MainPkg:      "./cmd/metric/",
			Dependencies: []string{"cmd", "pkg"},
		},
	}
	buildCmd      string
	buildLogLevel string
)

func init() {
	flag.StringVar(&buildCmd, "cmd", "build", "Build command (build, clean)")
	flag.StringVar(&buildLogLevel, "logLevel", "info", "Build command log level (debug, info, warn, error)")
}

func main() {
	flag.Parse()
	buildmgr.InitBuildManagerLog(&buildmgr.LogOption{
		Level:     buildmgr.ParseLogLevel(buildLogLevel, slog.LevelInfo),
		AddSource: false,
	})

	be, err := buildmgr.NewBuildEnv()
	if err != nil {
		panic(err)
	}

	slog.Debug("Build Option", slog.String("cmd", buildCmd))
	if buildCmd == "cleanDist" {
		err = buildmgr.CleanDist(be)
		if err != nil {
			slog.Error("Clean dist command error", slog.Any("err", err))
		} else {
			slog.Info("Clean dist command was completed with success")
		}
	} else {
		err = buildmgr.BuildAll(be, projects)
		if err != nil {
			slog.Error("Projects build error", slog.Any("err", err))
		} else {
			slog.Info("Projects building was completed with success")
		}
	}
}
