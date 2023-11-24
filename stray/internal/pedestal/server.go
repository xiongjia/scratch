package pedestal

import (
	"context"
	"sync"

	"google.golang.org/grpc"
)

type Pedestal struct {
	status pedestalStatus

	rpcSvc        NetworkService
	rpcSvcHandles []RpcServiceRegistrar
	// TODO REST Services and handlers
}

type RpcServiceRegistrar func(grpc.ServiceRegistrar)

func NewServer() (s *Pedestal, err error) {
	s = &Pedestal{
		status: stopped,
	}
	return s, nil
}

func (pd *Pedestal) AddRpcServiceHandle(handler RpcServiceRegistrar) {
	pd.rpcSvcHandles = append(pd.rpcSvcHandles, handler)
}

func (pd *Pedestal) Start(ctx context.Context) (err error) {
	var wg sync.WaitGroup

	rpcServ := newRPCServer(pd.rpcSvc, pd.rpcSvcHandles)
	rpcServ.start(&wg)

	wg.Wait()
	return nil
}
