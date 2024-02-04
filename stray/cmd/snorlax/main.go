package main

import (
	"log/slog"
	"stray/pkg/bulbasaur"
	"stray/pkg/dugtrio"
)

func main() {
	logger := dugtrio.NewSLog(dugtrio.SLogOptions{
		SLogBaseOptions: dugtrio.SLogBaseOptions{
			Level:     slog.LevelDebug,
			AddSource: true,
		},
	})
	dugtrio.SetDefaultLogger(logger)

	bulbasaur.NewEngine(bulbasaur.EngineOptions{
		Logger: logger,
	})
	logger.Debug("Test ended")
}
