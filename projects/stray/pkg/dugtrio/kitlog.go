package dugtrio

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

	logger *slog.Logger
	ctx    context.Context
}

func convertToSlogLevel(lvl string) slog.Level {
	switch lvl {
	case kitlevel.DebugValue().String():
		return slog.LevelDebug
	case kitlevel.WarnValue().String():
		return slog.LevelWarn
	case kitlevel.InfoValue().String():
		return slog.LevelInfo
	default:
		return slog.LevelError
	}
}

func (h *loggerHandler) Log(keyVals ...interface{}) error {
	srcSize := len(keyVals)
	lvl := ""
	attrs := make([]slog.Attr, srcSize/2+1)
	attrIdx := 0
	for i := 0; i < srcSize; i += 2 {
		k := keyVals[i]
		var v interface{} = errMissingValue
		if i+1 < srcSize {
			v = keyVals[i+1]
		}
		if k == kitlevel.Key() {
			lvl = fmt.Sprintf("%s", v)
			continue
		}
		attrs[attrIdx] = slog.Any(fmt.Sprint(k), v)
		attrIdx += 1
	}
	h.logger.LogAttrs(h.ctx, convertToSlogLevel(lvl), "", attrs...)
	return nil
}

func NewKitLoggerAdapterSlog(l *slog.Logger) kitlog.Logger {
	if l == nil {
		return kitlog.NewNopLogger()
	}
	handler := &loggerHandler{logger: l, ctx: context.Background()}
	logger := kitlevel.NewFilter(handler, kitlevel.AllowAll())
	return logger
}

func NewKitLoggerAdapterNop() kitlog.Logger {
	return kitlog.NewNopLogger()
}
