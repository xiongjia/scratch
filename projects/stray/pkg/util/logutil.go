package util

import (
	"log/slog"
	"os"
	"strings"
)

type (
	LogOption struct {
		Level     slog.Level
		AddSource bool
	}
)

func InitDefaultLog(opts *LogOption) {
	slog.SetDefault(NewLog(opts))
}

func NewLog(opts *LogOption) *slog.Logger {
	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level:     opts.Level,
		AddSource: opts.AddSource,
	})
	return slog.New(handler)
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
