package main

import (
	"log/slog"
	"stray/pkg/bulbasaur"
	"stray/pkg/dugtrio"
)

func main() {
	slog.SetDefault(dugtrio.NewSLog(dugtrio.SLogOptions{
		SLogBaseOptions: dugtrio.SLogBaseOptions{
			Level:     slog.LevelDebug,
			AddSource: true,
		},
	}))
	slog.Debug("debug test")

	bulbasaur.NewEngine()

	slog.Debug("Test ended")
}
