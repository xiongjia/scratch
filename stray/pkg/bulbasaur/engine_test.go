package bulbasaur_test

import (
	"log/slog"
	"stray/pkg/bulbasaur"
	"stray/pkg/dugtrio"
	"testing"
)

func TestEngine(t *testing.T) {
	logger := dugtrio.NewSlogForTesting(t, dugtrio.SLogBaseOptions{Level: slog.LevelDebug})

	engine, err := bulbasaur.NewEngine(bulbasaur.EngineOptions{Logger: logger})
	if err != nil {
		return
	}
	engine.Test()
}
