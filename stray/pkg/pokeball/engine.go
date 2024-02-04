package pokeball

import (
	"fmt"
	"log/slog"
	"net/http"
)

type EngineOptions struct {
	Logger *slog.Logger
}

type HandlerFunc func(http.ResponseWriter, *http.Request)

type Engine struct {
	logger *slog.Logger
	router map[string]HandlerFunc
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	key := req.Method + "-" + req.URL.Path
	if handler, ok := engine.router[key]; ok {
		handler(w, req)
	} else {
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
	}
}

func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	key := method + "-" + pattern
	engine.logger.Debug("Adding Route",
		slog.String("method", method), slog.String("pattern", pattern))
	engine.router[key] = handler
}

func (engine *Engine) Get(pattern string, handler HandlerFunc) {
	engine.addRoute("GET", pattern, handler)
}

func New(opts EngineOptions) *Engine {
	logger := opts.Logger
	if logger == nil {
		logger = slog.Default()
	}
	logger.Debug("creating a new pokeball engine")
	return &Engine{router: make(map[string]HandlerFunc), logger: logger}
}
