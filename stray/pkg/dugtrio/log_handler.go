package dugtrio

import (
	"fmt"
	"io"
	"log/slog"
	"os"
	"sync/atomic"
	"testing"

	"gopkg.in/natefinch/lumberjack.v2"
)

const (
	SLogFileMinSize    int        = 10
	SLogFileMinBackups int        = 3
	SLogDefaultLevel   slog.Level = slog.LevelInfo
)

var (
	defaultLogger atomic.Value
)

func init() {
	defaultLogger.Store(NewSLog(SLogOptions{
		SLogBaseOptions: SLogBaseOptions{
			Level:     SLogDefaultLevel,
			AddSource: false,
		},
		DisableStdout: false,
		LogFilename:   "",
	}))
}

func DefaultLogger() *slog.Logger {
	return defaultLogger.Load().(*slog.Logger)
}

func SetDefaultLogger(l *slog.Logger) {
	defaultLogger.Store(l)
}

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

func makeSLogHandlerOptions(opts *SLogBaseOptions) *slog.HandlerOptions {
	logOpts := &slog.HandlerOptions{Level: opts.Level, AddSource: opts.AddSource}
	if !logOpts.AddSource || !HasProjectRoot() {
		return logOpts
	}
	logOpts.ReplaceAttr = func(groups []string, a slog.Attr) slog.Attr {
		if a.Key == slog.SourceKey {
			source := a.Value.Any().(*slog.Source)
			source.File = convertSourceFilename(source.File)
		}
		return a
	}
	return logOpts
}

func NewSLog(opts SLogOptions) *slog.Logger {
	return slog.New(slog.NewJSONHandler(makeLogWriter(&opts),
		makeSLogHandlerOptions(&opts.SLogBaseOptions)))
}

func NewSLogNop() *slog.Logger {
	return slog.New(slog.NewTextHandler(&logNop{}, nil))
}

func NewSlogCallback(opts SLogBaseOptions, cb OnAppendLog) *slog.Logger {
	return slog.New(slog.NewJSONHandler(&logJsonCallback{callback: cb},
		makeSLogHandlerOptions(&opts)))
}

func NewSlogForTesting(t *testing.T, opts SLogBaseOptions) *slog.Logger {
	if t == nil {
		return NewSLogNop()
	}
	return NewSlogCallback(opts, func(logLine string) {
		if len(logLine) == 0 {
			return
		}
		t.Log(logLine)
		fmt.Println(logLine)
	})
}
