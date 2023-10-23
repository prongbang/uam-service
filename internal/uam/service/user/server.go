package user

import (
	"context"
	"github.com/prongbang/uam-service/internal/localizations"
	"github.com/prongbang/uam-service/internal/uam/service/role"
	"github.com/prongbang/uam-service/pkg/code"
	"github.com/prongbang/uam-service/pkg/core"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
)

type server struct {
	UserUc UseCase
	RoleUc role.UseCase
}

func (s *server) GetById(ctx context.Context, request *UserIdRequest) (*UserResponse, error) {
	payload := core.GrpcPayload(request.GetToken())

	id := request.GetId()
	if !core.IsUuid(&id) {
		return nil, status.New(codes.InvalidArgument, core.TranslateCtx(ctx, localizations.CommonInvalidData)).Err()
	}

	data := s.UserUc.GetById(ParamsGetById{ID: id, UserID: payload.Sub})
	if !core.IsUuid(data.ID) {
		return nil, status.New(codes.NotFound, core.TranslateCtx(ctx, localizations.CommonNotFoundData)).Err()
	}

	return &UserResponse{
		Code:    code.StatusOK,
		Message: http.StatusText(http.StatusOK),
		Data:    ToUserDataMapper(data),
	}, nil
}

func (s *server) Create(ctx context.Context, request *UserCreateRequest) (*UserResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *server) Update(ctx context.Context, request *UserUpdateRequest) (*UserResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *server) UpdatePassword(ctx context.Context, request *UserUpdatePasswordRequest) (*UserResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *server) UpdatePasswordMe(ctx context.Context, request *UserUpdatePasswordMeRequest) (*UserResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *server) Delete(ctx context.Context, request *UserDeleteRequest) (*UserDeleteResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *server) GetMe(ctx context.Context, request *UserMeRequest) (*UserResponse, error) {
	payload := core.GrpcPayload(request.GetToken())

	data := s.UserUc.GetById(ParamsGetById{ID: payload.Sub, UserID: payload.Sub})
	if !core.IsUuid(data.ID) {
		return nil, status.New(codes.NotFound, core.TranslateCtx(ctx, localizations.CommonNotFoundData)).Err()
	}

	return &UserResponse{
		Code:    code.StatusOK,
		Message: http.StatusText(http.StatusOK),
		Data:    ToUserDataMapper(data),
	}, nil
}

func (s *server) GetList(ctx context.Context, request *UserListRequest) (*UserListResponse, error) {
	payload := core.GrpcPayload(request.GetToken())
	paging := core.PagingRequest{
		Page:  int(request.GetPage()),
		Limit: int(request.GetLimit()),
	}
	if paging.Invalid() {
		return nil, status.New(codes.InvalidArgument, core.TranslateCtx(ctx, localizations.CommonInvalidData)).Err()
	}

	params := Params{
		UserID: payload.Sub,
	}

	getCount := func() int64 { return s.UserUc.Count(params) }

	getData := func(limit int, offset int) []User {
		params.LimitNo = paging.Limit
		params.OffsetNo = offset
		return s.UserUc.GetList(params)
	}
	resp := core.Pagination[User](paging.Page, paging.Limit, getCount, getData)

	return &UserListResponse{
		Code:    code.StatusOK,
		Message: http.StatusText(http.StatusOK),
		Data:    ToUserPagingMapper(resp),
	}, nil
}

func (s *server) mustEmbedUnimplementedUserServer() {}

func NewServer(
	userUc UseCase,
	roleUc role.UseCase,
) UserServer {
	return &server{
		UserUc: userUc,
		RoleUc: roleUc,
	}
}
