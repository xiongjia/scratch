package metric

import (
	"errors"
	"testing"

	"github.com/go-kit/log/level"
)

type (
	logHandler struct {
		t *testing.T
	}
)

func (h *logHandler) Append(entry *LogEntry) {
	if entry == nil {
		return
	}
	h.t.Log("entry", entry)
	h.t.Log("entry msg", entry.MergeMessage())
}

func TestKitLogAdapter(t *testing.T) {
	logAdapter := NewLogAdapter(&logHandler{t: t})
	_ = level.Debug(logAdapter).Log("msg", "test",
		"val1Num", 1, "val2Str", "str", "err", errors.New("test error"))
}
