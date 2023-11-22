package pedestal

import "context"

type Pedestal struct {
	status pedestalStatus

	rpcSvc  NetworkService
	restSvc NetworkService
}

func NewServer() (s *Pedestal, err error) {
	s = &Pedestal{
		status: stopped,
	}

	return s, nil
}

func (pd *Pedestal) Start(ctx context.Context) (err error) {

	return nil
}
