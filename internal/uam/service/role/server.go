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
	b := CreateRole{
		Name:  request.GetName(),
		Level: request.GetLevel(),
	}

	if b.Name == "" || b.Level < Level1 {
		return nil, status.New(codes.InvalidArgument, core.TranslateCtx(ctx, localizations.CommonInvalidData)).Err()
	}

	if err := s.RoleUc.Add(&b); err != nil {
		return nil, status.New(codes.InvalidArgument, core.TranslateCtx(ctx, err.Error())).Err()
	}

	return &RoleResponse{
		Code:    code.StatusOK,
		Message: http.StatusText(http.StatusOK),
		Data:    &RoleData{Id: *b.ID, Name: b.Name},
	}, nil
}

func (s server) Update(ctx context.Context, request *RoleUpdateRequest) (*RoleResponse, error) {
	id := request.GetId()
	b := UpdateRole{
		ID:    id,
		Name:  request.GetName(),
		Level: request.GetLevel(),
	}
	if core.IsNotUuid(&b.ID) || b.Name == "" || (b.Level == Level1 || b.Level < 0) {
		return nil, status.New(codes.InvalidArgument, core.TranslateCtx(ctx, localizations.CommonInvalidData)).Err()
	}

	if err := s.RoleUc.Update(&b); err != nil {
		return nil, status.New(codes.InvalidArgument, core.TranslateCtx(ctx, err.Error())).Err()
	}

	return &RoleResponse{
		Code:    code.StatusOK,
		Message: http.StatusText(http.StatusOK),
		Data:    &RoleData{Id: b.ID, Name: b.Name},
	}, nil
}

func (s server) Delete(ctx context.Context, request *RoleIdRequest) (*RoleDeleteResponse, error) {
	id := request.GetId()
	if core.IsNotUuid(&id) {
		return nil, status.New(codes.InvalidArgument, core.TranslateCtx(ctx, localizations.CommonInvalidData)).Err()
	}

	if err := s.RoleUc.Delete(id); err != nil {
		return nil, status.New(codes.InvalidArgument, core.TranslateCtx(ctx, err.Error())).Err()
	}

	return &RoleDeleteResponse{
		Code:    code.StatusOK,
		Message: http.StatusText(http.StatusOK),
	}, nil
}

func (s server) GetById(ctx context.Context, request *RoleIdRequest) (*RoleResponse, error) {
	payload := core.GrpcPayload(request.GetToken())

	if !core.IsUuid(&request.Id) {
		return nil, status.New(codes.InvalidArgument, core.TranslateCtx(ctx, localizations.CommonInvalidData)).Err()
	}

	params := ParamsGetById{ID: request.GetId(), UserID: payload.Sub}
	data := s.RoleUc.GetById(params)
	if data.ID == "" {
		return nil, status.New(codes.NotFound, core.TranslateCtx(ctx, localizations.CommonNotFoundData)).Err()
	}
	return &RoleResponse{
		Code:    code.StatusOK,
		Message: http.StatusText(http.StatusOK),
		Data:    &RoleData{Id: data.ID, Name: data.Name},
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
