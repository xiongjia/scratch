package main

import (
	"log/slog"
	"stray/internal/util"
)

func main() {
	util.InitSLog(&util.SLogOptions{Level: slog.LevelDebug, AddSource: true})

	slog.Debug("debug test")
}
