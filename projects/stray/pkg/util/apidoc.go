package util

import (
	"fmt"
	"net/http"
	swaggerUI "stray/third_party/swagger-ui"
)

func ApiDocUtilBind(mux *http.ServeMux, apiDoc any, router string) {
	mux.HandleFunc(fmt.Sprintf("GET /%s/openapi.yaml", router),
		HttpUtilYamlHandler(&apiDoc))
	mux.HandleFunc(fmt.Sprintf("GET /%s/{path...}", router),
		HttpUtilFsHandler(http.FS(swaggerUI.UIRoot), router))
}
