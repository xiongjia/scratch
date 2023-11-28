package pedestal

import (
	"fmt"
	"net"
	"net/http"
	"stray/internal/log"
	"strconv"
	"sync"
)

type RestfulServer struct {
	svc NetworkService
}

func newRestfulServer(svc NetworkService) *RestfulServer {
	return &RestfulServer{
		svc: svc,
	}
}

func (s *RestfulServer) start(wg *sync.WaitGroup) (err error) {
	mux := http.NewServeMux()
	mux.HandleFunc("/demo", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "test")
	})

	log.Infof("Restful server %s:%d", s.svc.IP, s.svc.Port)
	srv := &http.Server{
		Addr: net.JoinHostPort(
			s.svc.IP,
			strconv.Itoa(int(s.svc.Port)),
		),
		Handler: mux,
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		srv.ListenAndServe()
	}()
	return nil
}
