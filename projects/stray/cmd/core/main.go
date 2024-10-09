package main

import (
	"fmt"
	"log/slog"
	"runtime/debug"
	"stray/internal/server"
	"stray/pkg/util"
)

func main() {
	if bi, ok := debug.ReadBuildInfo(); ok {
		fmt.Printf("%+v\n", bi)
	}

	util.InitDefaultLog(&util.LogOption{
		Level:     slog.LevelDebug,
		AddSource: false,
	})

	serv, err := server.NewServer(server.ServerConfig{
		EnableApiDoc: true,
	})
	if err != nil {
		slog.Error("Server init error", slog.Any("err", err))
		return
	}

	wg := util.NewWaitGroup()
	wg.Go(func() {
		slog.Info("API Server Started (0.0.0.0:8897)")
		err = server.StartServer("0.0.0.0", 8897, serv)
		if err != nil {
			slog.Error("Server error", slog.Any("err", err))
		}
	})
	wg.Wait()
}
