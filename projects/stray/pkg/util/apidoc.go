package util

import (
	"fmt"
	"net/http"
	swaggerUI "stray/third_party/swagger-ui"
)

func ApiDocUtilBind(mux *http.ServeMux, apiDoc any, routePrefix string) {
	mux.HandleFunc(fmt.Sprintf("GET /%s/openapi.yaml", routePrefix),
		HttpUtilYamlHandler(&apiDoc))
	mux.HandleFunc(fmt.Sprintf("GET /%s/{path...}", routePrefix),
		HttpUtilFsHandler(http.FS(swaggerUI.UIRoot), routePrefix))
}
