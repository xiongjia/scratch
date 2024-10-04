package server

import (
	"encoding/json"
	"net/http"
	api "stray/internal/api/v1/server"
)

func (Server) GetVersion(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(api.VersionInfo{Major: VERSION_MAJOR, Minor: VERSION_MINOR, Patch: VERSION_PATCH})
}
