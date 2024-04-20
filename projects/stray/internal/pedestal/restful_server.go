package pedestal

import (
	"net"
	"net/http"
	"path/filepath"
	"stray/internal/log"
	"strconv"
	"sync"
	"time"

	"github.com/go-chi/chi/v5"

	assetfs "github.com/elazarl/go-bindata-assetfs"
	"github.com/go-chi/chi/middleware"
)

type RestfulServer struct {
	svc      NetworkService
	handlers []RestfulHandleRegistrar
	mux      *chi.Mux
}

type RestfulHandleRegistrar func(mux *chi.Mux)

func assertFsHandle(w http.ResponseWriter, r *http.Request,
	root *assetfs.AssetFS, path string) error {
	f, err := root.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	http.ServeContent(w, r, filepath.Base(path), time.Time{}, f)
	return nil
}

func newRestfulServer(svc NetworkService, handles []RestfulHandleRegistrar) *RestfulServer {
	serv := &RestfulServer{
		svc: svc,
		mux: chi.NewRouter(),
	}
	serv.handlers = append(serv.handlers, swaggerUiHandle)
	serv.handlers = append(serv.handlers, handles...)
	return serv
}

func (s *RestfulServer) start(wg *sync.WaitGroup) (err error) {
	s.mux.Use(middleware.Logger)
	for _, h := range s.handlers {
		h(s.mux)
	}

	// DEBUG MUX methods and middlewares
	chi.Walk(s.mux, func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		log.Infof("[%s]: '%s' has %d middlewares", method, route, len(middlewares))
		return nil
	})

	srv := &http.Server{
		Addr: net.JoinHostPort(
			s.svc.IP,
			strconv.Itoa(int(s.svc.Port)),
		),
		Handler: s.mux,
	}

	wg.Add(1)
	go func() {
		defer wg.Done()

		log.Infof("Restful server %s:%d", s.svc.IP, s.svc.Port)
		srv.ListenAndServe()
	}()
	return nil
}
