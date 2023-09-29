package grpc

import "github.com/prongbang/user-service/internal/user"

type GRPC interface {
	Register()
}

type gRPC struct {
	UserListener user.Listener
}

func (g *gRPC) Register() {
	g.UserListener.Serve()
}

func NewGRPC(userListener user.Listener) GRPC {
	return &gRPC{
		UserListener: userListener,
	}
}
