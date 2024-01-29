package util_test

import (
	"fmt"
	"log/slog"
	"os"
	"stray/pkg/dugtrio/util"
	"testing"

	kitlog "github.com/go-kit/log"
	"github.com/go-kit/log/level"
)

func TestKitLog(t *testing.T) {
	logOpts := &slog.HandlerOptions{
		Level:     slog.LevelDebug,
		AddSource: true,
	}
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, logOpts)))

	logger := kitlog.NewJSONLogger(kitlog.NewSyncWriter(os.Stdout))
	level.Debug(logger).Log("data", 1, "msg", "this message is at the debug level")
	l := level.ParseDefault("debug", level.ErrorValue())
	fmt.Printf("v = %v\n", l.String())

	logger2 := util.NewKitLoggerAdapter()
	level.Debug(logger2).Log("data", 1, "msg", "this message is at the debug level")

	slog.Error("test", "data", 1, "data", 2)
}
