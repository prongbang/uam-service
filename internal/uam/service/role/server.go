package role

import (
	"context"
	"github.com/prongbang/uam-service/internal/localizations"
	"github.com/prongbang/uam-service/pkg/code"
	"github.com/prongbang/uam-service/pkg/core"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
)

type server struct {
	RoleUc UseCase
}

func (s server) GetList(ctx context.Context, request *RoleListRequest) (*RoleListResponse, error) {
	payload := core.GrpcPayload(request.GetToken())

	params := Params{
		UserID: payload.Sub,
	}
	data := s.RoleUc.GetList(params)

	return &RoleListResponse{
		Code:    code.StatusOK,
		Message: http.StatusText(http.StatusOK),
		Data:    ToRoleListMapper(data),
	}, nil
}

func (s server) Add(ctx context.Context, request *RoleCreateRequest) (*RoleResponse, error) {
	//TODO implement me
	return nil, status.New(codes.Unimplemented, "Unimplemented").Err()
}

func (s server) Update(ctx context.Context, request *RoleUpdateRequest) (*RoleResponse, error) {
	//TODO implement me
	return nil, status.New(codes.Unimplemented, "Unimplemented").Err()
}

func (s server) Delete(ctx context.Context, request *RoleIdRequest) (*RoleResponse, error) {
	//TODO implement me
	return nil, status.New(codes.Unimplemented, "Unimplemented").Err()
}

func (s server) GetById(ctx context.Context, request *RoleIdRequest) (*RoleIdResponse, error) {
	payload := core.GrpcPayload(request.GetToken())

	if !core.IsUuid(&request.Id) {
		return nil, status.New(codes.InvalidArgument, core.TranslateCtx(ctx, localizations.CommonInvalidData)).Err()
	}

	params := ParamsGetById{ID: request.GetId(), UserID: payload.Sub}
	data := s.RoleUc.GetById(params)
	if data.ID == "" {
		return nil, status.New(codes.NotFound, core.TranslateCtx(ctx, localizations.CommonNotFoundData)).Err()
	}
	return &RoleIdResponse{
		Code:    code.StatusOK,
		Message: http.StatusText(http.StatusOK),
		Data:    &RoleResponse{Id: data.ID, Name: data.Name},
	}, nil
}

func (s server) mustEmbedUnimplementedRoleServer() {}

func NewServer(
	roleUc UseCase,
) RoleServer {
	return &server{
		RoleUc: roleUc,
	}
}
