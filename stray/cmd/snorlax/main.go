package main

import (
	"log/slog"
	"stray/internal/util"
	"stray/pkg/bulbasaur"
)

func main() {
	slog.SetDefault(util.MakeSLog(&util.SLogOptions{Level: slog.LevelDebug, AddSource: true}))
	slog.Debug("debug test")

	bulbasaur.NewEngine()

	slog.Debug("Test ended")
}
