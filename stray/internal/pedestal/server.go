package pedestal

import (
	"context"
	"sync"

	"google.golang.org/grpc"
)

type Pedestal struct {
	status pedestalStatus

	RpcSvc        NetworkService
	rpcSvcHandles []RpcServiceRegistrar

	RestfulSvc     NetworkService
	restfulHandles []RestfulHandleRegistrar
}

type RpcServiceRegistrar func(grpc.ServiceRegistrar)

func NewServer() (s *Pedestal, err error) {
	s = &Pedestal{
		status: stopped,
		RpcSvc: NetworkService{
			IP:   "0.0.0.0",
			Port: 8081,
		},
		RestfulSvc: NetworkService{
			IP:   "0.0.0.0",
			Port: 8082,
		},
	}
	return s, nil
}

func (pd *Pedestal) AddRpcServiceHandle(handler RpcServiceRegistrar) {
	pd.rpcSvcHandles = append(pd.rpcSvcHandles, handler)
}

func (pd *Pedestal) AddRestfulHandler(handler RestfulHandleRegistrar) {
	pd.restfulHandles = append(pd.restfulHandles, handler)
}

func (pd *Pedestal) Start(ctx context.Context) (err error) {
	var wg sync.WaitGroup

	rpcServ := newRPCServer(pd.RpcSvc, pd.rpcSvcHandles)
	rpcServ.start(&wg)

	restfulServ := newRestfulServer(pd.RestfulSvc, pd.restfulHandles)
	restfulServ.start(&wg)

	wg.Wait()
	return nil
}
