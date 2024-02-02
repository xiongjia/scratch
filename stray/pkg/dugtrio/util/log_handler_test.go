package util_test

import (
	"encoding/json"
	"log/slog"
	"stray/pkg/dugtrio/util"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type logTestItem struct {
	Msg   string `json:"msg,omitempty"`
	Num   int    `json:"num,omitempty"`
	Str   string `json:"string,omitempty"`
	Level string `json:"level,omitempty"`
}

func TestLogCallback(t *testing.T) {
	l := util.NewSlogCallback(util.SLogBaseOptions{Level: slog.LevelDebug}, func(logLine string) {
		if logLine == "" {
			return
		}
		var item logTestItem
		err := json.Unmarshal([]byte(logLine), &item)
		assert.NoError(t, err)
		assert.Equal(t, "Test Msg", item.Msg)
		assert.Equal(t, 2, item.Num)
		assert.Equal(t, "abc", item.Str)
		assert.True(t, strings.EqualFold("debug", item.Level))
	})
	l.Debug("Test Msg", "num", 2, "string", "abc")
}
