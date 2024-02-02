package util_test

import (
	"encoding/json"
	"log/slog"
	"stray/pkg/dugtrio/util"
	"strings"
	"testing"

	kitlevel "github.com/go-kit/log/level"
	"github.com/stretchr/testify/assert"
)

type kitLogItem struct {
	Time  string `json:"time,omitempty"`
	Level string `json:"level,omitempty"`
	Str   string `json:"string,omitempty"`
	Num   int    `json:"num,omitempty"`
}

func TestKitLog(t *testing.T) {
	l := util.NewSlogCallback(util.SLogBaseOptions{Level: slog.LevelDebug}, func(logLine string) {
		if logLine == "" {
			return
		}
		var item kitLogItem
		err := json.Unmarshal([]byte(logLine), &item)
		assert.NoError(t, err)
		assert.Equal(t, "this message is at the debug level", item.Str)
		assert.Equal(t, 1, item.Num)
		assert.True(t, strings.EqualFold("debug", item.Level))
	})
	kitlog := util.NewKitLoggerAdapterSlog(l)
	kitlevel.Debug(kitlog).Log("num", 1, "string", "this message is at the debug level")
}
