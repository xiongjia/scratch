package dugtrio_test

import (
	"encoding/json"
	"log/slog"
	"stray/pkg/dugtrio"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type logTestItemSrc struct {
	SrcFunc string `json:"function,omitempty"`
	SrcFile string `json:"file,omitempty"`
	SrcLine int    `json:"line,omitempty"`
}

type logTestItem struct {
	Source logTestItemSrc `json:"source"`

	Msg   string `json:"msg,omitempty"`
	Num   int    `json:"num,omitempty"`
	Str   string `json:"string,omitempty"`
	Level string `json:"level,omitempty"`
}

func TestLogCallback(t *testing.T) {
	opts := dugtrio.SLogBaseOptions{Level: slog.LevelDebug, AddSource: true}
	l := dugtrio.NewSlogCallback(opts, func(logLine string) {
		if logLine == "" {
			t.Logf("log Line is empty")
			return
		}
		t.Logf("log Line: %s", logLine)
		var item logTestItem
		err := json.Unmarshal([]byte(logLine), &item)
		assert.NoError(t, err)
		t.Logf("Log Item = %v", item)
		assert.Equal(t, "Test Msg", item.Msg)
		assert.Equal(t, 2, item.Num)
		assert.Equal(t, "abc", item.Str)
		assert.True(t, strings.EqualFold("debug", item.Level))
		assert.True(t, strings.Contains(item.Source.SrcFunc, "TestLogCallback"))
		assert.True(t, strings.Contains(item.Source.SrcFile, "log_handler_test.go"))
	})
	l.Debug("Test Msg", "num", 2, "string", "abc")
}
