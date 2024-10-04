package main

import (
	"log/slog"
	"stray/internal/server"
	"stray/pkg/utility"
)

func main() {
	utility.InitDefaultLog(&utility.LogOption{
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

	wg := utility.MakeWaitGroup()
	wg.Go(func() {
		err = server.StartServer("0.0.0.0", 8897, serv)
		if err != nil {
			slog.Error("Server error", slog.Any("err", err))
		}
	})
	wg.Wait()
}
