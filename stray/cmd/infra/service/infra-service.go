package service

import (
	"context"
	"net"
	"stray/gen/api/infra/v1"
	"stray/internal/log"
	"stray/internal/pedestal"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

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
	b, _ := apidoc.Asset("infra.swagger.yaml")
	pedestal.AddRestfulBuffHandle(mux, "/swagger-ui/openapi.yaml", b)
}

func grpcGatewayHandle(mux *chi.Mux, rpcSvc pedestal.NetworkService) {
	cc, err := grpc.DialContext(context.Background(), net.JoinHostPort(
		rpcSvc.IP,
		strconv.Itoa(int(rpcSvc.Port)),
	), grpc.WithBlock(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to dial server: %s", err.Error())
	}

	gwmux := runtime.NewServeMux()
	err = infra.RegisterGreeterHandlerClient(context.Background(), gwmux, infra.NewGreeterClient(cc))
	if err != nil {
		log.Fatalf("Failed to handle RPC Gateway: %s", err.Error())
	}
	mux.Handle("/*", gwmux)
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
	s.AddRestfulHandler(func(mux *chi.Mux) {
		grpcGatewayHandle(mux, s.RpcSvc)
	})
	s.Start(context.Background())
}
