package main

import (
	"log/slog"
	"net/http"
	"stray/cmd/snorlax/api"
	"stray/pkg/dugtrio"

	"github.com/go-chi/chi/v5"
)

func main() {
	logger := dugtrio.NewSLog(dugtrio.SLogOptions{
		SLogBaseOptions: dugtrio.SLogBaseOptions{
			Level:     slog.LevelDebug,
			AddSource: true,
		},
	})
	dugtrio.SetDefaultLogger(logger)

	// create a type that satisfies the `api.ServerInterface`, which contains an implementation of every operation from the generated code
	server := api.NewServer()

	r := chi.NewMux()

	// get an `http.Handler` that we can use
	h := api.HandlerFromMux(server, r)

	s := &http.Server{
		Handler: h,
		Addr:    "0.0.0.0:8080",
	}

	s.ListenAndServe()
}
