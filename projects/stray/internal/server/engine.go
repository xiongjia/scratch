package server

import (
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"strconv"

	api "stray/internal/api/v1/server"
	"stray/pkg/collection"
	"stray/pkg/util"
)

type (
	Server struct {
		engineMetric *EngineMetric
	}

	ServerConfig struct {
		EnableApiDoc bool
	}
)

var (
	_ api.ServerInterface = (*Server)(nil)
)

func makeServerMux(cfg *ServerConfig) (*http.ServeMux, error) {
	apiDoc, err := api.GetSwagger()
	if err != nil {
		return nil, err
	}
	mux := http.NewServeMux()
	if cfg.EnableApiDoc {
		slog.Debug("server api doc is enabled")
		util.ApiDocUtilBind(mux, apiDoc, PREFIX_API_DOC)
	}
	return mux, nil
}

func StartServer(host string, port int, srv http.Handler) error {
	addr := net.JoinHostPort(host, strconv.Itoa(port))
	httpServer := &http.Server{Handler: srv, Addr: addr}
	return httpServer.ListenAndServe()
}

func NewServer(cfg ...ServerConfig) (http.Handler, error) {
	servCfg := collection.FirstOrEmpty[ServerConfig](cfg)
	slog.Debug("Creating Server", slog.Any("cfg", servCfg))

	mux, err := makeServerMux(&servCfg)
	if err != nil {
		slog.Error("create server mux error", slog.Any("err", err), slog.Any("cfg", servCfg))
		return nil, err
	}
	engineMetric := NewEngineMetric(mux)
	server := &Server{engineMetric: engineMetric}
	handler := api.HandlerFromMuxWithBaseURL(server, mux, fmt.Sprintf("/%s", PREFIX_API_ROUTER))
	return util.HttpMiddlewareLog(handler), nil
}
