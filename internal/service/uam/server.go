package uam

import (
	context "context"
	"errors"
	"github.com/prongbang/user-service/internal/shared/user"
	"log"
)

// Server is used to implement uam.UserServer
type gRPCServer struct {
	Uc user.UseCase
	user.UnimplementedUserServer
}

func (a *gRPCServer) GetUser(ctx context.Context, request *user.UserRequest) (*user.UserResponse, error) {
	log.Printf("Received: username=%v", request.GetUsername())

	// Mock find uam by username
	if request.GetUsername() == "admin" {
		return &user.UserResponse{
			Id:       "1",
			Name:     "Administrator",
			Username: "admin",
			Password: "1234",
		}, nil
	}

	return nil, errors.New("404")
}

func NewServer(uc user.UseCase) user.UserServer {
	return &gRPCServer{
		Uc: uc,
	}
}
