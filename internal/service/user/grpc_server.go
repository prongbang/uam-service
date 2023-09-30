package user

import (
	context "context"
	"errors"
	"log"
)

// Server is used to implement user.UserServer
type gRPCServer struct {
	Uc UseCase
	UnimplementedUserServer
}

func (a *gRPCServer) GetUser(ctx context.Context, request *UserRequest) (*UserResponse, error) {
	log.Printf("Received: username=%v", request.GetUsername())

	// Mock find user by username
	if request.GetUsername() == "admin" {
		return &UserResponse{
			Id:       "1",
			Name:     "Administrator",
			Username: "admin",
			Password: "1234",
		}, nil
	}

	return nil, errors.New("404")
}

func NewServer(uc UseCase) UserServer {
	return &gRPCServer{
		Uc: uc,
	}
}
