package uam

type GRPC interface {
	Register()
}

type gRPC struct {
	GRPCListeners Listeners
}

func (g *gRPC) Register() {
	g.GRPCListeners.Serve()
}

func NewGRPC(gRPCListeners Listeners) GRPC {
	return &gRPC{
		GRPCListeners: gRPCListeners,
	}
}
