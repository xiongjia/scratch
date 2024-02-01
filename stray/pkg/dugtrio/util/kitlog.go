package util

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	kitlog "github.com/go-kit/log"
	kitlevel "github.com/go-kit/log/level"
)

var (
	errMissingValue = errors.New("(MISSING)")
)

type loggerHandler struct {
	kitlog.Logger
}

func parseKitLogLevel(v any) kitlevel.Value {
	return kitlevel.ParseDefault(fmt.Sprintf("%s", v), kitlevel.ErrorValue())
}

func convertToSlogLevel(lvl kitlevel.Value) slog.Level {
	return slog.LevelError
}

func appendKitLog(lvl kitlevel.Value, attrs []slog.Attr) {
	slog.LogAttrs(context.Background(), convertToSlogLevel(lvl), "", attrs...)
}

func (h *loggerHandler) Log(keyVals ...interface{}) error {
	srcSize := len(keyVals)
	lvl := kitlevel.ErrorValue()
	attrs := make([]slog.Attr, srcSize/2+1)
	attrIdx := 0
	for i := 0; i < srcSize; i += 2 {
		k := keyVals[i]
		var v interface{} = errMissingValue
		if i+1 < srcSize {
			v = keyVals[i+1]
		}
		if k == kitlevel.Key() {
			lvl = parseKitLogLevel(v)
			continue
		}
		attrs[attrIdx] = slog.Any(fmt.Sprint(k), v)
		attrIdx += 1
	}
	appendKitLog(lvl, attrs)
	return nil
}

func NewKitLoggerAdapter() kitlog.Logger {
	return kitlevel.NewFilter(&loggerHandler{}, kitlevel.AllowAll())
}

func NewKitLoggerAdapterNop() kitlog.Logger {
	return kitlog.NewNopLogger()
}
