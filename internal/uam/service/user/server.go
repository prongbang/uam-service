package user

import (
	"context"
	"github.com/prongbang/uam-service/internal/localizations"
	"github.com/prongbang/uam-service/internal/uam/service/role"
	"github.com/prongbang/uam-service/pkg/code"
	"github.com/prongbang/uam-service/pkg/common"
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

	params := ParamsGetById{
		ID:      id,
		Payload: payload,
	}
	data := s.UserUc.GetById(params)
	if !core.IsUuid(data.ID) {
		return nil, status.New(codes.NotFound, core.TranslateCtx(ctx, localizations.CommonNotFoundData)).Err()
	}

	return &UserResponse{
		Code:    code.StatusOK,
		Message: http.StatusText(http.StatusOK),
		Data:    FromUserMapper(data),
	}, nil
}

func (s *server) Add(ctx context.Context, request *UserCreateRequest) (*UserResponse, error) {
	payload := core.GrpcPayload(request.GetToken())

	if request.Username == "" && request.Email == "" {
		return nil, status.New(codes.InvalidArgument, core.TranslateCtx(ctx, localizations.CommonInvalidData)).Err()
	}

	if (request.Username != "" && len(request.Username) < UsernameMin) || (request.Email != "" && !common.IsEmail(request.Email)) || len(request.Password) < PasswordMin {
		return nil, status.New(codes.InvalidArgument, core.TranslateCtx(ctx, localizations.CommonInvalidData)).Err()
	}

	body := CreateUser{
		Username:  request.Username,
		Password:  request.Password,
		Email:     request.Email,
		FirstName: request.FirstName,
		LastName:  request.LastName,
		Avatar:    request.Avatar,
		Mobile:    request.Mobile,
		CreatedBy: payload.UserID,
		Payload:   payload,
	}
	usr, err := s.UserUc.Add(&body)
	if err != nil {
		return nil, status.New(codes.InvalidArgument, core.TranslateCtx(ctx, (*err).Message)).Err()
	}

	return &UserResponse{
		Code:    code.StatusOK,
		Message: http.StatusText(http.StatusOK),
		Data:    FromUserMapper(usr),
	}, nil
}

func (s *server) Update(ctx context.Context, request *UserUpdateRequest) (*UserResponse, error) {
	payload := core.GrpcPayload(request.GetToken())

	if (request.Username != "" && len(request.Username) < UsernameMin) || (request.Email != "" && !common.IsEmail(request.Email)) {
		return nil, status.New(codes.InvalidArgument, core.TranslateCtx(ctx, localizations.CommonInvalidData)).Err()
	}

	body := UpdateUser{
		ID:        request.Id,
		Username:  request.Username,
		Email:     request.Email,
		FirstName: request.FirstName,
		LastName:  request.LastName,
		Avatar:    request.Avatar,
		Mobile:    request.Mobile,
		Payload:   payload,
	}

	usr, err := s.UserUc.Update(&body)
	if err != nil {
		return nil, status.New(codes.InvalidArgument, core.TranslateCtx(ctx, (*err).Message)).Err()
	}

	return &UserResponse{
		Code:    code.StatusOK,
		Message: http.StatusText(http.StatusOK),
		Data:    FromUserMapper(usr),
	}, nil
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
	payload := core.GrpcPayload(request.GetToken())

	id := request.GetId()
	if core.IsNotUuid(&id) {
		return nil, status.New(codes.InvalidArgument, core.TranslateCtx(ctx, localizations.CommonInvalidData)).Err()
	}

	// Check user under
	body := DeleteUser{ID: id, Payload: payload}
	if err := s.UserUc.Delete(body); err != nil {
		return nil, status.New(codes.InvalidArgument, core.TranslateCtx(ctx, localizations.CommonInvalidData)).Err()
	}
	return &UserDeleteResponse{
		Code:    code.StatusOK,
		Message: http.StatusText(http.StatusOK),
	}, nil
}

func (s *server) GetMe(ctx context.Context, request *UserMeRequest) (*UserResponse, error) {
	payload := core.GrpcPayload(request.GetToken())

	params := ParamsGetById{ID: payload.UserID, Payload: payload}
	data := s.UserUc.GetById(params)
	if !core.IsUuid(data.ID) {
		return nil, status.New(codes.NotFound, core.TranslateCtx(ctx, localizations.CommonNotFoundData)).Err()
	}

	return &UserResponse{
		Code:    code.StatusOK,
		Message: http.StatusText(http.StatusOK),
		Data:    FromUserMapper(data),
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
		Payload: payload,
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
		Data:    FromUserPagingMapper(resp),
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
