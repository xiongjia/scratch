package buildmgr

import (
	"errors"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

type (
	BuildEnv struct {
		Verbose      bool
		ProjectDir   string
		GoCmd        string
		BuildScripts []string
		OutputDir    string
		BuildDebug   bool
		BuildGOOS    string
		RuntimeGOOS  string
	}
)

func NewBuildEnv() (buildEnv *BuildEnv, err error) {
	wrkDir, err := os.Getwd()
	if err != nil {
		slog.Error("Get work dir error", slog.Any("err", err))
		return nil, err
	}
	buildEnv = &BuildEnv{
		Verbose:      true,
		ProjectDir:   ".",
		GoCmd:        "go",
		BuildDebug:   true,
		BuildGOOS:    os.Getenv("GOOS"),
		RuntimeGOOS:  runtime.GOOS,
		OutputDir:    filenameClean(filepath.Join(wrkDir, "./dist")),
		BuildScripts: []string{"build.go"},
	}
	slog.Debug("New build env is created",
		slog.Any("env", buildEnv), slog.String("wrkDir", wrkDir))
	return buildEnv, nil
}

func (be *BuildEnv) GetGOOS() string {
	goos := be.BuildGOOS
	if len(goos) == 0 {
		goos = be.RuntimeGOOS
	}
	return goos
}

func (be *BuildEnv) MakeBinFilename(filename string) string {
	if !strings.EqualFold(be.GetGOOS(), "windows") {
		return filenameClean(filepath.Join(be.OutputDir, filename))
	} else {
		return filenameClean(filepath.Join(be.OutputDir, fmt.Sprintf("%s.exe", filename)))
	}
}

func (be *BuildEnv) MakeProjectFullFilename(filename string) string {
	if filepath.IsAbs(filename) {
		return filenameClean(filename)
	}
	return filenameClean(filepath.Join(be.ProjectDir, filename))
}

func (be *BuildEnv) ShouldRebuildAssets(target string, srcDirs ...string) bool {
	info, err := os.Stat(target)
	if err != nil {
		return true
	}
	currentBuild := info.ModTime()
	srcAreNewer := false
	for _, buildScript := range be.BuildScripts {
		scriptInfo, err := os.Stat(be.MakeProjectFullFilename(buildScript))
		if err != nil {
			return true
		}
		if scriptInfo.ModTime().After(currentBuild) {
			return true
		}
	}
	for _, srcDir := range srcDirs {
		fullSrcDir := be.MakeProjectFullFilename(srcDir)
		slog.Debug("Checking dependency", slog.String("src", fullSrcDir))
		stop := errors.New("no need to iterate further")
		filepath.Walk(fullSrcDir, func(path string, info os.FileInfo, err error) error {
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
