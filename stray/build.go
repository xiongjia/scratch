package main

import (
	"flag"
	"fmt"
	"log/slog"
	"stray/pkg/ditto"
)

func makeProject(be *ditto.BuildEnv) *ditto.Project {
	ldflags := fmt.Sprintf("-X stray/pkg/dugtrio.ProjectRoot=\"%s\"", be.ProjectDir)
	dependencies := []string{"cmd", "internal", "pkg"}
	return &ditto.Project{
		OutputBinaries: []ditto.Binary{
			{Output: "pikachu",
				MainPkg:      "cmd/pikachu/main.go",
				LDFlags:      ldflags,
				Dependencies: dependencies},
			{Output: "snorlax",
				MainPkg:      "cmd/snorlax/main.go",
				LDFlags:      ldflags,
				Dependencies: dependencies},
		},
	}
}

func main() {
	be, err := ditto.NewBuildEnv()
	if err != nil {
		panic(err)
	}
	be.Log.Info("Build Script is running", slog.String("env", be.DumpEnv()))
	project := makeProject(be)
	be.Log.Debug("Build Project", slog.Any("project", project))
	cmd := flag.Arg(0)
	if len(cmd) == 0 {
		cmd = "buildAll"
	}
	be.RunBuildCommand(cmd, project)
}
