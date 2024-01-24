package util

import (
	"io"
	"log/slog"
	"os"

	"gopkg.in/natefinch/lumberjack.v2"
)

type SLogOptions struct {
	Level         slog.Level
	AddSource     bool
	LogFilename   string
	DisableStdout bool
}

func makeFsLogWriter(opts *SLogOptions) io.Writer {
	return &lumberjack.Logger{
		Filename:   opts.LogFilename,
		MaxSize:    500,
		MaxBackups: 3,
		Compress:   true,
	}
}

func makeLogWriter(opts *SLogOptions) io.Writer {
	if len(opts.LogFilename) == 0 && !opts.DisableStdout {
		return os.Stdout
	} else if len(opts.LogFilename) != 0 && !opts.DisableStdout {
		return io.MultiWriter(makeFsLogWriter(opts), os.Stdout)
	} else if len(opts.LogFilename) != 0 && opts.DisableStdout {
		return makeFsLogWriter(opts)
	} else {
		return makeDummyWriter()
	}
}

func MakeSLog(opts *SLogOptions) *slog.Logger {
	logOpts := &slog.HandlerOptions{
		Level:     opts.Level,
		AddSource: opts.AddSource,
	}
	return slog.New(slog.NewJSONHandler(makeLogWriter(opts), logOpts))
}
