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
	"net/http"
)

type server struct {
	Uc     user.UseCase
	AuthUc UseCase
}

func (u *server) RestEnforce(ctx context.Context, request *AuthEnforceRequest) (*AuthEnforceResponse, error) {
	if u.AuthUc.RestEnforce(request.Subject, request.Object, request.Action) {
		return &AuthEnforceResponse{}, nil
	}
	return nil, status.New(codes.PermissionDenied, core.TranslateCtx(ctx, localizations.CommonPermissionDenied)).Err()
}

func (u *server) RbacEnforce(ctx context.Context, request *AuthEnforceRequest) (*AuthEnforceResponse, error) {
	if u.AuthUc.RbacEnforce(request.Subject, request.Object, request.Action) {
		return &AuthEnforceResponse{}, nil
	}
	return nil, status.New(codes.PermissionDenied, core.TranslateCtx(ctx, localizations.CommonPermissionDenied)).Err()
}

func (u *server) VerifyToken(ctx context.Context, request *AuthVerifyTokenRequest) (*AuthVerifyTokenResponse, error) {
	accessToken := request.GetToken()

	if err := u.AuthUc.VerifyToken(accessToken); err != nil {
		return nil, status.New(codes.Unauthenticated, core.TranslateCtx(ctx, localizations.CommonUnauthenticated)).Err()
	}
	return &AuthVerifyTokenResponse{
		Code:    code.StatusOK,
		Message: http.StatusText(http.StatusOK),
	}, nil
}

func (u *server) Login(ctx context.Context, request *AuthRequest) (*AuthResponse, error) {
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
	return &AuthResponse{
		Code: code.StatusOK,
		Data: &AuthCredential{Token: credential.Token, Roles: credential.Roles},
	}, nil
}

func (u *server) mustEmbedUnimplementedAuthServer() {}

func NewServer(
	uc user.UseCase,
	authUc UseCase,
) AuthServer {
	return &server{
		Uc:     uc,
		AuthUc: authUc,
	}
}
