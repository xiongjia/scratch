package server

import (
	"log/slog"
	"net"
	"net/http"
	"strconv"

	api "stray/internal/api/v1/server"
	"stray/pkg/collection"
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
	handler := api.HandlerFromMuxWithBaseURL(&Server{}, mux, PREFIX_API_ROUTER)
	return handler, nil
}
