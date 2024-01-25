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
	Verbose    bool   `json:"verbose"`
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
		Verbose:     true,
		ProjectDir:  ".",
		GoCmd:       "go",
		BuildDebug:  true,
		BuildGOOS:   os.Getenv("GOOS"),
		RuntimeGOOS: runtime.GOOS,
		OutputDir:   "./build-output",
	}
	project = Project{
		OutputBinaries: []Binary{
			{Output: "pikachu", MainPkg: "cmd/pikachu/main.go"},
		},
	}
)

func initLog() {
	logOpts := &slog.HandlerOptions{}
	if buildEnv.Verbose {
		logOpts.Level = slog.LevelDebug
	}
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, logOpts)))
}

func main() {
	flag.BoolVar(&buildEnv.Verbose, "verbose", false, "Print verbose build log")
	flag.Parse()
	initLog()

	slog.Info("Staring Build script ", slog.Any("env", buildEnv), slog.Any("project", project))
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
		rmFiles(makeOutputBinFilename(binary.Output))
	}
}

func buildAll() {
	for _, binary := range project.OutputBinaries {
		buildBinary(binary)
	}
}

func makeOutputBinFilename(filename string) string {
	goos := buildEnv.BuildGOOS
	if len(goos) == 0 {
		goos = buildEnv.RuntimeGOOS
	}
	if !strings.EqualFold(goos, "windows") {
		return filepath.Join(buildEnv.OutputDir, filename)
	}
	return filepath.Join(buildEnv.OutputDir, fmt.Sprintf("%s.exe", filename))
}

func buildBinary(bin Binary) {
	slog.Debug("Build binary: ", "bin", bin)

	output := makeOutputBinFilename(bin.Output)
	binPkg := filepath.Join(buildEnv.ProjectDir, bin.MainPkg)
	shouldRebuild := shouldRebuildAssets(output, binPkg,
		filepath.Join(buildEnv.ProjectDir, "cmd"),
		filepath.Join(buildEnv.ProjectDir, "internal"),
		filepath.Join(buildEnv.ProjectDir, "pkg"))
	if !shouldRebuild {
		slog.Info("No source code change. (SKIP)", slog.Any("bin", bin))
		return
	}

	args := []string{"build"}
	if buildEnv.Verbose {
		args = append(args, "-v")
	}
	if buildEnv.BuildDebug {
		args = append(args, "-gcflags", "all=-N -l")
	} else {
		args = append(args, "-trimpath")
	}
	args = append(args, "-o", output, binPkg)
	exitCode, err := invokeCmd(".", buildEnv.GoCmd, args...)
	if err != nil {
		slog.Error("go build command error", slog.Any("error", err))
		return
	}
	if exitCode != 0 {
		slog.Error("go build command error exit code != 0", slog.Int("exitCode", exitCode))
		return
	}
	slog.Info("Build process finished", slog.Any("pkg", bin))
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

func shouldRebuildAssets(target string, srcDirs ...string) bool {
	info, err := os.Stat(target)
	if err != nil {
		// If the file doesn't exist, we must rebuild it
		return true
	}
	currentBuild := info.ModTime()
	srcAreNewer := false
	stop := errors.New("no need to iterate further")
	for _, srcDir := range srcDirs {
		filepath.Walk(srcDir, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.ModTime().After(currentBuild) {
				srcAreNewer = true
				return stop
			}
			return nil
		})
	}
	return srcAreNewer
}
