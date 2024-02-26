package main

import (
	"flag"
	"fmt"
	"log/slog"
	"stray/pkg/ditto"
)

func main() {
	be, err := ditto.NewBuildEnv()
	if err != nil {
		panic(err)
	}
	be.Log.Info("Build Script is running", slog.String("env", be.DumpEnv()))
	ldflags := fmt.Sprintf("-X stray/pkg/dugtrio.ProjectRoot=\"%s\"", be.ProjectDir)
	project := &ditto.Project{
		OutputBinaries: []ditto.Binary{
			{Output: "pikachu", MainPkg: "cmd/pikachu/main.go", LDFlags: ldflags},
			{Output: "snorlax", MainPkg: "cmd/snorlax/main.go", LDFlags: ldflags},
		},
	}
	be.Log.Debug("Build Project", slog.Any("project", project))
	cmd := flag.Arg(0)
	if len(cmd) == 0 {
		cmd = "buildAll"
	}
	be.RunBuildCommand(cmd, project)
}
