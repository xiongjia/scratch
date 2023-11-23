package service

import (
	"context"
	"stray/gen/api/infra/v1"
	"stray/internal/log"
	"stray/internal/pedestal"

	"google.golang.org/grpc"
)

type infraServ struct {
	infra.UnimplementedGreeterServer
}

func (s *infraServ) SayHello(ctx context.Context, in *infra.HelloRequest) (*infra.HelloReply, error) {
	log.Debugf("Main service request")
	return &infra.HelloReply{Message: in.Name + " infra module "}, nil
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
	s.Start(context.Background())
}
