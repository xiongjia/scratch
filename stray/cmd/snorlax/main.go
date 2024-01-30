package main

import (
	"log/slog"
	"stray/pkg/bulbasaur"
	"stray/pkg/dugtrio/util"
)

func main() {
	slog.SetDefault(util.NewSLog(&util.SLogOptions{Level: slog.LevelDebug, AddSource: true}))
	slog.Debug("debug test")

	bulbasaur.NewEngine()

	slog.Debug("Test ended")
}
