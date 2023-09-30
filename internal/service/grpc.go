package service

import "github.com/prongbang/user-service/internal/service/user"

type GRPC interface {
	Register()
}

type gRPC struct {
	UserListener user.GRPCListener
}

func (g *gRPC) Register() {
	g.UserListener.Serve()
}

func NewGRPC(userListener user.GRPCListener) GRPC {
	return &gRPC{
		UserListener: userListener,
	}
}
