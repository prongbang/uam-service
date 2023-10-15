package uam

import (
	"context"
	"errors"
	"github.com/prongbang/uam-service/internal/localizations"
	"github.com/prongbang/uam-service/internal/shared/auth"
	"github.com/prongbang/uam-service/internal/shared/user"
	"github.com/prongbang/uam-service/pkg/code"
	"github.com/prongbang/uam-service/pkg/common"
	"github.com/prongbang/uam-service/pkg/core"
)

// Server is used to implement uam.UamServer
type gRPCServer struct {
	Uc     user.UseCase
	AuthUc auth.UseCase
	UnimplementedUamServer
}

func (u *gRPCServer) Login(ctx context.Context, request *LoginRequest) (*LoginResponse, error) {
	username := request.GetUsername()
	email := request.GetEmail()
	password := request.GetPassword()

	if password == "" || (email == "" && username == "") {
		return nil, errors.New(core.TranslateCtx(ctx, localizations.CommonInvalidData))
	}

	if email != "" && username != "" {
		if !common.IsEmail(email) {
			return nil, errors.New(core.TranslateCtx(ctx, localizations.CommonInvalidData))
		}
	}

	req := auth.Login{
		Username: username,
		Email:    email,
		Password: password,
	}
	credential, err := u.AuthUc.Login(req)
	if err != nil {
		return nil, errors.New(core.TranslateCtx(ctx, err.Error()))
	}
	return &LoginResponse{
		Code: code.StatusOK,
		Data: &Credential{Token: credential.Token, Roles: credential.Roles},
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
