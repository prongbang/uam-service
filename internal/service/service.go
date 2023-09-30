package service

type Service interface {
	NewAPI() API
	NewGRPC() GRPC
}

type service struct {
	API  API
	GRPC GRPC
}

func (s *service) NewAPI() API {
	return s.API
}

func (s *service) NewGRPC() GRPC {
	return s.GRPC
}

func NewService(api API, gRPC GRPC) Service {
	return &service{
		API:  api,
		GRPC: gRPC,
	}
}
