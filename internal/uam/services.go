package uam

type Services interface {
	NewAPI() API
	NewGRPC() GRPC
}

type services struct {
	API  API
	GRPC GRPC
}

func (s *services) NewAPI() API {
	return s.API
}

func (s *services) NewGRPC() GRPC {
	return s.GRPC
}

func NewService(api API, gRPC GRPC) Services {
	return &services{
		API:  api,
		GRPC: gRPC,
	}
}
