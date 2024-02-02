package util

import (
	"io"
	"log/slog"
	"os"

	"gopkg.in/natefinch/lumberjack.v2"
)

const (
	SLogFileMinSize    int = 10
	SLogFileMinBackups int = 3
)

type logNop struct{ io.Writer }

func (l *logNop) Write(p []byte) (n int, err error) { return len(p), nil }

type OnAppendLog func(logLine string)

type logJsonCallback struct {
	io.Writer
	callback OnAppendLog
}

func (l *logJsonCallback) Write(p []byte) (n int, err error) {
	l.callback(string(p))
	return len(p), nil
}

type SLogBaseOptions struct {
	Level     slog.Level
	AddSource bool
}

type SLogOptions struct {
	SLogBaseOptions

	DisableStdout         bool
	LogFilename           string
	LogFileMaxSize        int
	LogFileMaxBackups     int
	LogFileBackupCompress bool
}

func makeFsLogWriter(opts *SLogOptions) io.Writer {
	return &lumberjack.Logger{
		Filename:   opts.LogFilename,
		MaxSize:    max(SLogFileMinSize, opts.LogFileMaxSize),
		MaxBackups: max(SLogFileMinBackups, opts.LogFileMaxBackups),
		Compress:   opts.LogFileBackupCompress,
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
		return &logNop{}
	}
}

func NewSLog(opts SLogOptions) *slog.Logger {
	logOpts := &slog.HandlerOptions{
		Level:     opts.Level,
		AddSource: opts.AddSource,
	}
	return slog.New(slog.NewJSONHandler(makeLogWriter(&opts), logOpts))
}

func NewSlogCallback(opts SLogBaseOptions, cb OnAppendLog) *slog.Logger {
	logOpts := &slog.HandlerOptions{
		Level:     opts.Level,
		AddSource: opts.AddSource,
	}
	return slog.New(slog.NewJSONHandler(&logJsonCallback{callback: cb}, logOpts))
}
