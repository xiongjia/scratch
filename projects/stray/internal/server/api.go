package server

import (
	"fmt"
	"log/slog"
	"net/http"
	api "stray/internal/api/v1/server"
	"stray/pkg/util"
)

// (GET /version)
func (Server) GetVersion(w http.ResponseWriter, r *http.Request) {
	util.HttpUtilWriteObjectQuite(w, http.StatusOK,
		&api.VersionInfo{Major: VERSION_MAJOR, Minor: VERSION_MINOR, Patch: VERSION_PATCH})
}

// (POST /v1/debug/echo)
func (Server) PostV1DebugEcho(w http.ResponseWriter, r *http.Request) {
	opts, err := util.JsonReaderToObject[api.DebugEchoOptions](r.Body)
	if err != nil {
		util.HttpUtilWriteError(w, http.StatusBadRequest, err)
		return
	}

	slog.Debug("debug echo opts", slog.Any("opts", opts))
	util.HttpUtilWriteObjectQuite(w, http.StatusOK, &api.DebugEchoResponse{
		Msg: fmt.Sprintf("input message = %s", opts.Msg),
	})
}
