package dugtrio_test

import (
	"encoding/json"
	"log/slog"
	"stray/pkg/dugtrio"
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
	l := dugtrio.NewSlogCallback(dugtrio.SLogBaseOptions{Level: slog.LevelDebug, AddSource: true}, func(logLine string) {
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
	kitlog := dugtrio.NewKitLoggerAdapterSlog(l)
	kitlevel.Debug(kitlog).Log("num", 1, "string", "this message is at the debug level")
}
