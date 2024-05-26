package main

import (
	"log/slog"
	"net/http"
	"os"
	api "pumpkin/api/mgrsrv"
	server "pumpkin/internal/server"
	"pumpkin/pkg/util"
	swaggerUI "pumpkin/third_party/swagger-ui"

	"github.com/ghodss/yaml"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	log := util.NewSLog(util.LoggerOption{
		Level:     slog.LevelDebug,
		AddSource: true,
	})

	log.Debug("Manager Server")

	swagger, err := api.GetSwagger()
	if err != nil {
		log.Error("Error loading swagger spec", slog.Any("err", err))
		os.Exit(1)
	}

	r := chi.NewMux()
	r.Use(middleware.Logger)

	r.Get("/swagger-ui/*", http.StripPrefix("/swagger-ui/", http.FileServer(http.FS(swaggerUI.UIRoot))).ServeHTTP)
	r.Get("/swagger-ui/openapi.yaml", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/x-yaml")
		data, _ := yaml.Marshal(&swagger)
		_, _ = w.Write(data)
		w.WriteHeader(http.StatusOK)
	})

	s := &http.Server{
		Handler: api.HandlerFromMux(server.NewMgrSrv(), r),
		Addr:    "0.0.0.0:8080",
	}
	s.ListenAndServe()
}
