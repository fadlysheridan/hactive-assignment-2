package transport

import "assignment-2/app/interface/container"

type Tp struct {
	Transport *tp
}

func SetupTransport(constainer container.Container) *Tp {
	transport := NewTransport(constainer.Usecase)
	return &Tp{
		Transport: transport,
	}
}