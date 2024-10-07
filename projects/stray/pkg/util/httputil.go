package util

import (
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/ghodss/yaml"
)

func HttpMiddlewareLog(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		handler.ServeHTTP(w, r)
		slog.Debug("HTTP Request",
			slog.String("method", r.Method),
			slog.String("path", r.URL.Path),
			slog.Duration("duration", time.Since(start)))
	})
}

func HttpUtilWriteObject[T any](w http.ResponseWriter, statusCode int, obj T) error {
	w.WriteHeader(statusCode)
	return JsonWriterFromObject(w, obj)
}

func HttpUtilWriteObjectQuite[T any](w http.ResponseWriter, statusCode int, obj T) {
	err := HttpUtilWriteObject(w, statusCode, obj)
	if err != nil {
		slog.Error("Write HTTP Response error",
			slog.Any("err", err), slog.Int("status", statusCode), slog.Any("response", obj))
	}
}

func HttpUtilWriteError(w http.ResponseWriter, statusCode int, errObj error) {
	http.Error(w, errObj.Error(), statusCode)
}

func HttpUtilYamlHandler(content any) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/x-yaml")
		data, _ := yaml.Marshal(content)
		_, _ = w.Write(data)
		w.WriteHeader(http.StatusOK)
	}
}

func HttpUtilFsHandler(root http.FileSystem, routePrefix string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		http.StripPrefix(fmt.Sprintf("/%s/", routePrefix), http.FileServer(root)).ServeHTTP(w, r)
	}
}
