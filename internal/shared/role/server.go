package role

import (
	"context"
	"errors"
	"github.com/prongbang/uam-service/internal/localizations"
	"github.com/prongbang/uam-service/internal/shared/user"
	"github.com/prongbang/uam-service/pkg/core"
)

type server struct {
	UserUc user.UseCase
	RoleUc UseCase
	UnimplementedRoleServer
}

func (s server) GetById(ctx context.Context, request *RoleIdRequest) (*RoleResponse, error) {
	if !core.IsUuid(&request.Id) {
		return nil, errors.New(core.TranslateCtx(ctx, localizations.CommonInvalidData))
	}

	data := s.RoleUc.GetById(request.GetId())
	if data.ID == "" {
		return nil, errors.New(core.TranslateCtx(ctx, localizations.CommonNotFoundData))
	}
	return &RoleResponse{
		Id:    data.ID,
		Name:  data.Name,
		Level: data.Level,
	}, nil
}

func NewServer(
	userUc user.UseCase,
	roleUc UseCase,
) RoleServer {
	return &server{
		UserUc: userUc,
		RoleUc: roleUc,
	}
}
