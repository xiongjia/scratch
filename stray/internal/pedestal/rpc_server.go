package pedestal

type RPCServer struct {
	svc NetworkService
}

func newRPCServer(svc NetworkService) *RPCServer {
	return &RPCServer{
		svc: svc,
	}
}

func (s *RPCServer) start() (err error) {

	return
}
