package auth

import (
	"context"
	"github.com/prongbang/uam-service/internal/localizations"
	"github.com/prongbang/uam-service/internal/uam/service/user"
	"github.com/prongbang/uam-service/pkg/code"
	"github.com/prongbang/uam-service/pkg/common"
	"github.com/prongbang/uam-service/pkg/core"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct {
	Uc     user.UseCase
	AuthUc UseCase
	UnimplementedAuthServer
}

func (u *server) Login(ctx context.Context, request *LoginRequest) (*LoginResponse, error) {
	username := request.GetUsername()
	email := request.GetEmail()
	password := request.GetPassword()

	stInvalid := status.New(codes.InvalidArgument, core.TranslateCtx(ctx, localizations.CommonInvalidData))

	if password == "" || (email == "" && username == "") {
		return nil, stInvalid.Err()
	}

	if email != "" && username != "" {
		if !common.IsEmail(email) {
			return nil, stInvalid.Err()
		}
	}

	req := Login{
		Username: username,
		Email:    email,
		Password: password,
	}
	credential, err := u.AuthUc.Login(req)
	if err != nil {
		return nil, status.New(codes.InvalidArgument, core.TranslateCtx(ctx, err.Error())).Err()
	}
	return &LoginResponse{
		Code: code.StatusOK,
		Data: &LoginCredential{Token: credential.Token, Roles: credential.Roles},
	}, nil
}

func NewServer(
	uc user.UseCase,
	authUc UseCase,
) AuthServer {
	return &server{
		Uc:     uc,
		AuthUc: authUc,
	}
}
