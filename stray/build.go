package main

import (
	"errors"
	"flag"
	"fmt"
	"log/slog"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

type BuildEnv struct {
	ProjectDir string `json:"project_dir"`
	GoCmd      string `json:"go_cmd"`

	BuildDebug  bool   `json:"build_debug"`
	BuildGOOS   string `json:"build_goos"`
	RuntimeGOOS string `json:"runtime_goos"`

	OutputDir string `json:"output_dir"`
}

type Project struct {
	OutputBinaries []Binary `json:"output_binaries"`
}

type Binary struct {
	Output       string   `json:"output"`
	MainPkg      string   `json:"main_pkg"`
	Dependencies []string `json:"dependencies"`
}

var (
	buildEnv = BuildEnv{
		ProjectDir:  ".",
		GoCmd:       "go",
		BuildDebug:  true,
		BuildGOOS:   os.Getenv("GOOS"),
		RuntimeGOOS: runtime.GOOS,
		OutputDir:   ".",
	}
	project = Project{
		OutputBinaries: []Binary{
			{Output: "stray-infra", MainPkg: "cmd/infra/main.go"},
			{Output: "stray-core", MainPkg: "cmd/core/main.go"},
			{Output: "stray-tool", MainPkg: "cmd/core/main.go"},
		},
	}
)

func init() {
	logOpts := &slog.HandlerOptions{Level: slog.LevelDebug}
	handler := slog.NewJSONHandler(os.Stdout, logOpts)
	logger := slog.New(handler)
	slog.SetDefault(logger)
}

func main() {
	flag.Parse()
	slog.Debug("Build Script ",
		slog.Any("env", buildEnv),
		slog.Any("project", project))

	cmd := flag.Arg(0)
	if len(cmd) == 0 {
		cmd = "buildAll"
	}
	runCommand(cmd)
}

func runCommand(cmd string) {
	switch cmd {
	case "buildAll":
		buildAll()
	case "cleanAll":
		cleanAll()
	default:
		slog.Error("Unknown command", "cmd", cmd)
		panic("Unknown command")
	}
}

func cleanAll() {
	slog.Debug("cleanAll")
	for _, binary := range project.OutputBinaries {
		rmFiles(filepath.Join(buildEnv.OutputDir, makeOsBinFilename(binary.Output)))
	}
}

func buildAll() {
	for _, binary := range project.OutputBinaries {
		buildBinary(binary)
	}
}

func makeOsBinFilename(filename string) string {
	goos := buildEnv.BuildGOOS
	if len(goos) == 0 {
		goos = buildEnv.RuntimeGOOS
	}
	if !strings.EqualFold(goos, "windows") {
		return filename
	}
	return fmt.Sprintf("%s.exe", filename)
}

func buildBinary(bin Binary) {
	slog.Debug("Build binary: ", "bin", bin)

	args := []string{"build", "-v"}
	if buildEnv.BuildDebug {
		args = append(args, "-gcflags", "all=-N -l")
	} else {
		args = append(args, "-trimpath")
	}
	args = append(args,
		"-o", filepath.Join(buildEnv.OutputDir, makeOsBinFilename(bin.Output)),
		filepath.Join(buildEnv.ProjectDir, bin.MainPkg))
	// TODO check exit code
	invokeCmd(".", buildEnv.GoCmd, args...)
}

func rmFiles(paths ...string) {
	for _, path := range paths {
		slog.Debug("remove file : ", "path", path)
		os.RemoveAll(path)
	}
}

func invokeCmd(wrkDir string, cmd string, args ...string) (int, error) {
	slog.Debug("Executing command",
		slog.String("cmd", cmd), slog.Any("args", args))

	ecmd := exec.Command(cmd, args...)
	ecmd.Stdout = os.Stdout
	ecmd.Stderr = os.Stderr
	ecmd.Dir = wrkDir
	err := ecmd.Run()
	if err != nil {
		slog.Error("exec command error", "err", err)
		return 0, err
	}
	exitCode := ecmd.ProcessState.ExitCode()
	slog.Debug("Build exec exit code ", "exit", exitCode)
	return exitCode, nil
}

func shouldRebuildAssets(target, srcdir string) bool {
	info, err := os.Stat(target)
	if err != nil {
		// If the file doesn't exist, we must rebuild it
		return true
	}

	// Check if any of the files in gui/ are newer than the asset file. If
	// so we should rebuild it.
	currentBuild := info.ModTime()
	assetsAreNewer := false
	stop := errors.New("no need to iterate further")
	filepath.Walk(srcdir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.ModTime().After(currentBuild) {
			assetsAreNewer = true
			return stop
		}
		return nil
	})

	return assetsAreNewer
}
