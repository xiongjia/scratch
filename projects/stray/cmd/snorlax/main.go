package main

import (
	"log/slog"
	"net/http"
	"os"
	"stray/cmd/snorlax/api"
	"stray/pkg/dugtrio"

	"github.com/ghodss/yaml"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	logger := dugtrio.NewSLog(dugtrio.SLogOptions{
		SLogBaseOptions: dugtrio.SLogBaseOptions{
			Level:     slog.LevelDebug,
			AddSource: true,
		},
	})
	dugtrio.SetDefaultLogger(logger)

	swagger, err := api.GetSwagger()
	if err != nil {
		logger.Error("Error loading swagger spec: %s", err)
		os.Exit(1)
	}

	// create a type that satisfies the `api.ServerInterface`, which contains an implementation of every operation from the generated code
	// server := api.NewServer()

	r := chi.NewMux()
	r.Use(middleware.Logger)
	// r.Use(middleware.OapiRequestValidator(swagger))

	r.Get("/api", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/x-yaml")
		data, _ := yaml.Marshal(&swagger)
		_, _ = w.Write(data)
		// w.WriteHeader(http.StatusOK)
	})

	// get an `http.Handler` that we can use
	// h := api.HandlerFromMux(server, r)

	s := &http.Server{
		Handler: r,
		Addr:    "0.0.0.0:8080",
	}

	s.ListenAndServe()
}
