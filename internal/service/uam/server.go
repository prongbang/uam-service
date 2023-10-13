package uam

import (
	"context"
	"errors"
	"github.com/prongbang/uam-service/internal/localizations"
	"github.com/prongbang/uam-service/internal/shared/auth"
	"github.com/prongbang/uam-service/internal/shared/user"
	"github.com/prongbang/uam-service/pkg/code"
	"github.com/prongbang/uam-service/pkg/common"
)

// Server is used to implement uam.UamServer
type gRPCServer struct {
	Uc     user.UseCase
	AuthUc auth.UseCase
	UnimplementedUamServer
}

func (u *gRPCServer) Login(context context.Context, request *LoginRequest) (*LoginResponse, error) {
	username := request.GetUsername()
	email := request.GetEmail()
	password := request.GetPassword()

	if password == "" || (email == "" && username == "") {
		return nil, errors.New(localizations.CommonInvalidData)
	}

	if email != "" && username != "" {
		if !common.IsEmail(email) {
			return nil, errors.New(localizations.CommonInvalidData)
		}
	}

	credential, err := u.AuthUc.Login(auth.Login{
		Username: username,
		Email:    email,
		Password: password,
	})
	if err != nil {
		return &LoginResponse{
			Code: code.StatusBadRequest,
		}, errors.New(localizations.Translate(localizations.En, err.Error()))
	}
	return &LoginResponse{
		Code: code.StatusOK,
		Data: &Credential{
			Token: credential.Token,
			Roles: credential.Roles,
		},
	}, nil
}

func NewServer(
	uc user.UseCase,
	authUc auth.UseCase,
) UamServer {
	return &gRPCServer{
		Uc:     uc,
		AuthUc: authUc,
	}
}
