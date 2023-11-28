package pedestal

import (
	"context"
	"log"
	"sync"

	"github.com/go-chi/chi/v5"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Pedestal struct {
	status pedestalStatus

	rpcSvc        NetworkService
	rpcSvcHandles []RpcServiceRegistrar

	restfulSvc     NetworkService
	restfulHandles []RestfulHandleRegistrar
}

type RpcServiceRegistrar func(grpc.ServiceRegistrar)

func NewServer() (s *Pedestal, err error) {
	s = &Pedestal{
		status: stopped,
		rpcSvc: NetworkService{
			IP:   "0.0.0.0",
			Port: 8081,
		},
		restfulSvc: NetworkService{
			IP:   "0.0.0.0",
			Port: 8082,
		},
	}
	return s, nil
}

func grpcGatewayHandle(mux *chi.Mux) {
	cc, err := grpc.DialContext(context.Background(), "0.0.0.0:8081", grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}

	gwmux := runtime.NewServeMux()

	mux.Handle("/", gwmux)

}

func (pd *Pedestal) AddRpcServiceHandle(handler RpcServiceRegistrar) {
	pd.rpcSvcHandles = append(pd.rpcSvcHandles, handler)
}

func (pd *Pedestal) AddRestfulHandler(handler RestfulHandleRegistrar) {
	pd.restfulHandles = append(pd.restfulHandles, handler)
}

func (pd *Pedestal) Start(ctx context.Context) (err error) {
	var wg sync.WaitGroup

	rpcServ := newRPCServer(pd.rpcSvc, pd.rpcSvcHandles)
	rpcServ.start(&wg)

	// err = hello.RegisterGreeterHandler(context.Background(), gwmux, cc)
	// if err != nil {
	//		log.Fatalln("Failed to gwmux:", err)
	// }

	pd.AddRestfulHandler(grpcGatewayHandle)
	restfulServ := newRestfulServer(pd.restfulSvc, pd.restfulHandles)
	restfulServ.start(&wg)

	wg.Wait()
	return nil
}
