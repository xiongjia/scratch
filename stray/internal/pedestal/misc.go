package pedestal

import (
	"net/http"
	docui "stray/gen/swagger"
	"strings"

	assetfs "github.com/elazarl/go-bindata-assetfs"
	"github.com/go-chi/chi/v5"
)

func AddServerFsHandle(mux *chi.Mux, path string, root http.FileSystem) {
	fserv := http.FileServer(root)
	mux.Get(path, func(w http.ResponseWriter, r *http.Request) {
		rctx := chi.RouteContext(r.Context())
		pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
		fs := http.StripPrefix(pathPrefix, fserv)
		fs.ServeHTTP(w, r)
	})
}

func swaggerUiHandle(mux *chi.Mux) {
	AddServerFsHandle(mux, "/swagger-ui/*", &assetfs.AssetFS{
		Asset:    docui.Asset,
		AssetDir: docui.AssetDir,
		Prefix:   "swagger-ui",
		Fallback: "index.html",
	})
}
