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

	data := s.RoleUc.GetListByUnderRoles(payload.Roles)

	list := []*RoleResponse{}
	for _, u := range data {
		list = append(list, &RoleResponse{Id: u.ID, Name: u.Name, Level: u.Level})
	}

	return &RoleListResponse{
		Code:    code.StatusOK,
		Message: http.StatusText(http.StatusOK),
		Data:    list,
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

func (s server) mustEmbedUnimplementedRoleServer() {}

func NewServer(
	roleUc UseCase,
) RoleServer {
	return &server{
		RoleUc: roleUc,
	}
}
