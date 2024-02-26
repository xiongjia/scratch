package ditto

import (
	"errors"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

type Binary struct {
	Output       string   `json:"output"`
	MainPkg      string   `json:"main_pkg"`
	Dependencies []string `json:"dependencies"`
}

var (
	BuildEnv = struct {
		Verbose    bool   `json:"verbose"`
		ProjectDir string `json:"project_dir"`
		GoCmd      string `json:"go_cmd"`

		BuildDebug  bool   `json:"build_debug"`
		BuildGOOS   string `json:"build_goos"`
		RuntimeGOOS string `json:"runtime_goos"`

		OutputDir string `json:"output_dir"`
	}{
		Verbose:     true,
		ProjectDir:  ".",
		GoCmd:       "go",
		BuildDebug:  true,
		BuildGOOS:   os.Getenv("GOOS"),
		RuntimeGOOS: runtime.GOOS,
		OutputDir:   "./build-output",
	}
)

func init() {
	wrkDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	BuildEnv.ProjectDir = strings.ReplaceAll(filepath.Clean(filepath.Clean(wrkDir)), "\\", "/")
}

func makeOutputBinFilename(filename string) string {
	goos := BuildEnv.BuildGOOS
	if len(goos) == 0 {
		goos = BuildEnv.RuntimeGOOS
	}
	if !strings.EqualFold(goos, "windows") {
		return filepath.Join(BuildEnv.OutputDir, filename)
	}
	return filepath.Join(BuildEnv.OutputDir, fmt.Sprintf("%s.exe", filename))
}

func rmFiles(paths ...string) {
	for _, path := range paths {
		slog.Debug("remove file : ", "path", path)
		os.RemoveAll(path)
	}
}

func shouldRebuildAssets(target string, srcDirs ...string) bool {
	info, err := os.Stat(target)
	if err != nil {
		return true
	}
	currentBuild := info.ModTime()

	buildScriptInfo, err := os.Stat(filepath.Join(BuildEnv.ProjectDir, "build.go"))
	if err != nil {
		return true
	}
	if buildScriptInfo.ModTime().After(currentBuild) {
		return true
	}
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
