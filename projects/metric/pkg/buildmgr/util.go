package buildmgr

import (
	"log/slog"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type (
	LogOption struct {
		Level     slog.Level
		AddSource bool
	}
)

func filenameClean(filename string) string {
	if len(filename) <= 0 {
		return filename
	}
	return filepath.ToSlash(filepath.Clean(filename))
}

func invokeBuildCmd(wrkDir string, cmd string, args ...string) (int, error) {
	slog.Debug("Executing command", slog.String("cmd", cmd), slog.Any("args", args))
	ecmd := exec.Command(cmd, args...)
	ecmd.Stdout = os.Stdout
	ecmd.Stderr = os.Stderr
	ecmd.Dir = wrkDir
	err := ecmd.Run()
	if err != nil {
		slog.Error("exec command error", slog.Any("err", err))
		return -1, err
	}
	exitCode := ecmd.ProcessState.ExitCode()
	slog.Debug("Build exec exit code ", slog.Int("exit", exitCode))
	return exitCode, nil
}

func ParseLogLevel(level string, defaultLevel slog.Level) slog.Level {
	if strings.EqualFold(level, "debug") {
		return slog.LevelDebug
	} else if strings.EqualFold(level, "info") {
		return slog.LevelInfo
	} else if strings.EqualFold(level, "warn") {
		return slog.LevelWarn
	} else if strings.EqualFold(level, "error") {
		return slog.LevelError
	} else {
		return defaultLevel
	}
}

func NewSLog(opts *LogOption) *slog.Logger {
	return slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level:     opts.Level,
		AddSource: opts.AddSource,
	}))
}

func InitBuildManagerLog(opts *LogOption) {
	slog.SetDefault(NewSLog(opts))
}
