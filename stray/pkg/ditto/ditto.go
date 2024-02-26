package ditto

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

type Project struct {
	OutputBinaries []Binary
}

type Binary struct {
	Output       string
	MainPkg      string
	Dependencies []string
	LDFlags      string
}

type BuildEnv struct {
	Verbose      bool
	ProjectDir   string
	GoCmd        string
	BuildScripts []string
	OutputDir    string
	BuildDebug   bool
	BuildGOOS    string
	RuntimeGOOS  string

	Log *slog.Logger
}

func NewBuildEnv() (buildEnv *BuildEnv, err error) {
	wrkDir, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	buildEnv = &BuildEnv{
		Verbose:     true,
		ProjectDir:  ".",
		GoCmd:       "go",
		BuildDebug:  true,
		BuildGOOS:   os.Getenv("GOOS"),
		RuntimeGOOS: runtime.GOOS,
		OutputDir:   "./build-output",
	}
	buildEnv.ProjectDir = projectFilenameClean(wrkDir)
	flag.BoolVar(&buildEnv.Verbose, "verbose", false, "Print verbose build log")
	flag.Parse()
	logOpts := &slog.HandlerOptions{}
	if buildEnv.Verbose {
		logOpts.Level = slog.LevelDebug
	}
	buildEnv.Log = slog.New(slog.NewJSONHandler(os.Stdout, logOpts))
	return buildEnv, nil
}

func (be *BuildEnv) rmFiles(paths ...string) {
	for _, path := range paths {
		be.Log.Debug("remove file : ", "path", path)
		os.RemoveAll(path)
	}
}

func (be *BuildEnv) invokeCmd(wrkDir string, cmd string, args ...string) (int, error) {
	be.Log.Debug("Executing command",
		slog.String("cmd", cmd), slog.Any("args", args))
	ecmd := exec.Command(cmd, args...)
	ecmd.Stdout = os.Stdout
	ecmd.Stderr = os.Stderr
	ecmd.Dir = wrkDir
	err := ecmd.Run()
	if err != nil {
		be.Log.Error("exec command error", "err", err)
		return 0, err
	}
	exitCode := ecmd.ProcessState.ExitCode()
	be.Log.Debug("Build exec exit code ", "exit", exitCode)
	return exitCode, nil
}

func (be *BuildEnv) DumpEnv() string {
	var b strings.Builder
	b.WriteString("ProjectDir: ")
	b.WriteString(be.ProjectDir)
	return b.String()
}

func (be *BuildEnv) ConvertToFullFilename(filename string) string {
	if filepath.IsAbs(filename) {
		return filename
	}
	return filepath.Join(be.ProjectDir, filename)
}

func (be *BuildEnv) ConvertToOutputBinFilename(filename string) string {
	goos := be.BuildGOOS
	if len(goos) == 0 {
		goos = be.RuntimeGOOS
	}
	if !strings.EqualFold(goos, "windows") {
		return projectFilenameClean(filepath.Join(be.OutputDir, filename))
	}
	return projectFilenameClean(filepath.Join(be.OutputDir, fmt.Sprintf("%s.exe", filename)))
}

func (be *BuildEnv) shouldRebuildAssets(target string, srcDirs ...string) bool {
	info, err := os.Stat(target)
	if err != nil {
		return true
	}

	currentBuild := info.ModTime()
	srcAreNewer := false
	stop := errors.New("no need to iterate further")

	for _, buildScript := range be.BuildScripts {
		scriptInfo, err := os.Stat(be.ConvertToFullFilename(buildScript))
		if err != nil {
			return true
		}
		if scriptInfo.ModTime().After(currentBuild) {
			return true
		}
	}
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

func (be *BuildEnv) RunBuildCommand(cmd string, p *Project) {
	switch cmd {
	case "buildAll":
		be.buildAll(p)
	case "cleanAll":
		be.cleanAll(p)
	default:
		be.Log.Error("Unknown command", slog.String("cmd", cmd))
		panic(fmt.Sprintf("Unknown command %s", cmd))
	}
}

func (be *BuildEnv) cleanAll(p *Project) {
	be.Log.Debug("cleanAll")
	for _, binary := range p.OutputBinaries {
		be.rmFiles(be.ConvertToOutputBinFilename(binary.Output))
	}
}

func (be *BuildEnv) buildAll(p *Project) {
	for _, binary := range p.OutputBinaries {
		be.buildBinary(binary)
	}
}

func (be *BuildEnv) buildBinary(bin Binary) {
	be.Log.Debug("Build binary: ", "bin", bin)
	output := be.ConvertToOutputBinFilename(bin.Output)
	binPkg := be.ConvertToFullFilename(bin.MainPkg)
	shouldRebuild := be.shouldRebuildAssets(output, binPkg,
		be.ConvertToFullFilename("cmd"),
		be.ConvertToFullFilename("internal"),
		be.ConvertToFullFilename("pkg"))
	if !shouldRebuild {
		be.Log.Info("No source code change. (SKIP)", slog.Any("bin", bin))
		return
	}

	args := []string{"build"}
	if be.Verbose {
		args = append(args, "-v")
	}
	if be.BuildDebug {
		args = append(args, "-gcflags", "all=-N -l")
	} else {
		args = append(args, "-trimpath")
	}

	if len(bin.LDFlags) > 0 {
		args = append(args, "-ldflags", bin.LDFlags)
	}
	args = append(args, "-o", output, binPkg)

	exitCode, err := be.invokeCmd(".", be.GoCmd, args...)
	if err != nil {
		be.Log.Error("go build command error", slog.Any("error", err))
		return
	}
	if exitCode != 0 {
		be.Log.Error("go build command error exit code != 0", slog.Int("exitCode", exitCode))
		return
	}
	be.Log.Info("Build process finished", slog.Any("pkg", bin))
}
