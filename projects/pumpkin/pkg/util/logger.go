package util

import (
	"log/slog"
	"os"
)

type LoggerOption struct {
	Level     slog.Level
	AddSource bool
}

func NewSLog(opts LoggerOption) *slog.Logger {
	jsonLog := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level:     opts.Level,
		AddSource: opts.AddSource,
	})
	return slog.New(jsonLog)
}
