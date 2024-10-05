package server

import (
	"log/slog"
	"net"
	"net/http"
	"strconv"

	api "stray/internal/api/v1/server"
	"stray/pkg/collection"
	swaggerUI "stray/third_party/swagger-ui"

	"github.com/ghodss/yaml"
)

type (
	Server struct {
	}

	ServerConfig struct {
		EnableApiDoc bool
	}
)

var (
	_ api.ServerInterface = (*Server)(nil)
)

func StartServer(host string, port int, srv http.Handler) error {
	addr := net.JoinHostPort(host, strconv.Itoa(port))
	httpServer := &http.Server{Handler: srv, Addr: addr}
	return httpServer.ListenAndServe()
}

func NewServer(cfg ...ServerConfig) (http.Handler, error) {
	servCfg := collection.FirstOrEmpty[ServerConfig](cfg)
	slog.Debug("Creating Server", slog.Any("cfg", servCfg))

	mux := http.NewServeMux()

	apiDoc, err := api.GetSwagger()
	if err != nil {
		return nil, err
	}
	mux.HandleFunc("GET /apidoc/openapi.yaml", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/x-yaml")
		data, _ := yaml.Marshal(&apiDoc)
		_, _ = w.Write(data)
		w.WriteHeader(http.StatusOK)
	})
	mux.HandleFunc("GET /apidoc/{path...}", func(w http.ResponseWriter, r *http.Request) {
		handle := http.FileServer(http.FS(swaggerUI.UIRoot))
		http.StripPrefix("/apidoc/", handle).ServeHTTP(w, r)
	})
	handler := api.HandlerFromMuxWithBaseURL(&Server{}, mux, "/api")
	return handler, nil
}
