package pedestal

import (
	"net"
	"strconv"

	"stray/internal/log"

	"google.golang.org/grpc"
)

type RPCServer struct {
	svc     NetworkService
	handles []RpcServiceRegistrar
}

func newRPCServer(svc NetworkService, handles []RpcServiceRegistrar) *RPCServer {
	return &RPCServer{
		svc:     svc,
		handles: handles,
	}
}

func (s *RPCServer) start() (err error) {
	var opts []grpc.ServerOption

	listen, err := net.Listen("tcp", net.JoinHostPort(
		s.svc.IP,
		strconv.Itoa(int(s.svc.Port)),
	))
	if err != nil {
		log.Fatalf("Failed to listen:", err)
		return err
	}

	// TODO gRPC server opts
	serv := grpc.NewServer(opts...)
	for _, h := range s.handles {
		h(serv)
	}
	go func() {
		log.Infof("RPC Server started %s:%d", s.svc.IP, s.svc.Port)
		serv.Serve(listen)
	}()
	return
}
