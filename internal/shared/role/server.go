package role

import (
	"context"
	"github.com/prongbang/uam-service/internal/localizations"
	"github.com/prongbang/uam-service/pkg/core"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct {
	RoleUc UseCase
	UnimplementedRoleServer
}

func (s server) GetById(ctx context.Context, request *RoleIdRequest) (*RoleResponse, error) {
	if !core.IsUuid(&request.Id) {
		return nil, status.New(codes.InvalidArgument, core.TranslateCtx(ctx, localizations.CommonInvalidData)).Err()
	}

	data := s.RoleUc.GetById(request.GetId())
	if data.ID == "" {
		return nil, status.New(codes.NotFound, core.TranslateCtx(ctx, localizations.CommonNotFoundData)).Err()
	}
	return &RoleResponse{
		Id:    data.ID,
		Name:  data.Name,
		Level: data.Level,
	}, nil
}

func NewServer(
	roleUc UseCase,
) RoleServer {
	return &server{
		RoleUc: roleUc,
	}
}
