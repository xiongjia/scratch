package service

import (
	"context"
	"net/http"
	"stray/gen/api/infra/v1"
	"stray/internal/log"
	"stray/internal/pedestal"

	"github.com/go-chi/chi/v5"
	"google.golang.org/grpc"

	apidoc "stray/gen/api_doc/infra/v1"
)

type infraServ struct {
	infra.UnimplementedGreeterServer
}

func (s *infraServ) SayHello(ctx context.Context, in *infra.HelloRequest) (*infra.HelloReply, error) {
	log.Debugf("Main service request")
	return &infra.HelloReply{Message: in.Name + " infra module "}, nil
}

func swaggerDoc(mux *chi.Mux) {
	log.Infof("restful register")
	mux.Get("/swagger-ui/openapi.yaml", func(w http.ResponseWriter, r *http.Request) {
		b, _ := apidoc.Asset("infra.swagger.yaml")
		w.Write(b)
	})
}

func MainService() {
	s, err := pedestal.NewServer()
	if err != nil {
		log.Panicf("Core service start error: %s", err.Error())
	}

	svc := &infraServ{}
	s.AddRpcServiceHandle(func(r grpc.ServiceRegistrar) {
		infra.RegisterGreeterServer(r, svc)
	})
	s.AddRestfulHandler(swaggerDoc)
	s.Start(context.Background())
}
