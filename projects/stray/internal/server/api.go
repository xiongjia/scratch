package server

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	api "stray/internal/api/v1/server"
	"stray/pkg/collection"
)

// (GET /version)
func (Server) GetVersion(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(api.VersionInfo{Major: VERSION_MAJOR, Minor: VERSION_MINOR, Patch: VERSION_PATCH})
}

// (POST /v1/debug/echo)
func (Server) PostV1DebugEcho(w http.ResponseWriter, r *http.Request) {
	opts, err := collection.JsonReaderToObject[api.DebugEchoOptions](r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	slog.Debug("debug echo opts", slog.Any("opts", opts))

	response := &api.DebugEchoResponse{
		Msg: fmt.Sprintf("input message = %s", opts.Msg),
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
