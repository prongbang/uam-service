package service

import "github.com/prongbang/user-service/internal/service/uam"

type GRPC interface {
	Register()
}

type gRPC struct {
	UserListener uam.GRPCListener
}

func (g *gRPC) Register() {
	g.UserListener.Serve()
}

func NewGRPC(userListener uam.GRPCListener) GRPC {
	return &gRPC{
		UserListener: userListener,
	}
}
