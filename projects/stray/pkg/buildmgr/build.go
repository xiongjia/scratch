package buildmgr

import (
	"log/slog"
	"os"
)

type (
	BinaryBuildOption struct {
		Output       string
		MainPkg      string
		Dependencies []string
		LDFlags      string
	}
)

func CleanDist(be *BuildEnv) error {
	slog.Info("Remove Dist dir", slog.String("dist", be.OutputDir))
	err := os.RemoveAll(be.OutputDir)
	if err != nil {
		slog.Error("Remove dist dir error",
			slog.Any("err", err), slog.String("dist", be.OutputDir))
	}
	return err
}

func BuildAll(be *BuildEnv, opts []*BinaryBuildOption) error {
	for _, opt := range opts {
		err := Build(be, opt)
		if err != nil {
			return err
		}
	}
	return nil
}

func Build(be *BuildEnv, opt *BinaryBuildOption) error {
	slog.Debug("Binary build option", slog.Any("opt", opt))
	output := be.MakeBinFilename(opt.Output)
	mainPkg := be.MakeProjectFullFilename(opt.MainPkg)

	shouldRebuild := be.ShouldRebuildAssets(output, opt.Dependencies...)
	if !shouldRebuild {
		slog.Info("No source code change. (SKIP)", slog.Any("bin", opt))
		return nil
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
	if len(opt.LDFlags) > 0 {
		args = append(args, "-ldflags", opt.LDFlags)
	}
	args = append(args, "-o", output, mainPkg)

	exitCode, err := invokeBuildCmd(".", be.GoCmd, args...)
	if err != nil {
		slog.Error("go build command error", slog.Any("error", err))
		return err
	}
	if exitCode != 0 {
		slog.Error("go build command error exit code != 0", slog.Int("exitCode", exitCode))
		return nil
	}
	return nil
}
