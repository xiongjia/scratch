package server

import (
	"encoding/json"
	"net/http"
	api "pumpkin/api/mgrsrv"
)

var _ api.ServerInterface = (*MgrSrv)(nil)

type MgrSrv struct{}

func NewMgrSrv() MgrSrv {
	return MgrSrv{}
}

func (MgrSrv) GetPing(w http.ResponseWriter, r *http.Request, params api.GetPingParams) {
	resp := api.Pong{
		Ping: "pong pong",
	}
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(resp)
}
