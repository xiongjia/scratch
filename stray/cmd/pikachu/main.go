package main

import (
	"log/slog"
	"stray/internal/util"
)

func main() {
	slog.SetDefault(util.MakeSLog(&util.SLogOptions{Level: slog.LevelDebug, AddSource: true}))
	slog.Debug("debug test")
}
