package user

import (
	"context"
	"github.com/prongbang/uam-service/internal/localizations"
	"github.com/prongbang/uam-service/internal/shared/role"
	"github.com/prongbang/uam-service/pkg/code"
	"github.com/prongbang/uam-service/pkg/core"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"net/http"
)

type server struct {
	UserUc UseCase
	RoleUc role.UseCase
}

func (s *server) GetMe(ctx context.Context, request *UserMeRequest) (*UserResponse, error) {
	payload := core.PayloadByToken(request.GetToken())

	data := s.UserUc.GetById(payload.Sub)
	if !core.IsUuid(data.ID) {
		return nil, status.New(codes.NotFound, core.TranslateCtx(ctx, localizations.CommonNotFoundData)).Err()
	}

	var lastLogin *timestamppb.Timestamp
	if data.LastLogin != nil {
		lastLogin = timestamppb.New(*data.LastLogin)
	}
	return &UserResponse{
		Code:    code.StatusOK,
		Message: http.StatusText(http.StatusOK),
		Data: &UserData{
			Id:        *data.ID,
			Username:  data.Username,
			Email:     data.Email,
			FirstName: data.FirstName,
			LastName:  data.LastName,
			Avatar:    data.Avatar,
			Mobile:    data.Mobile,
			Flag:      int32(data.Flag),
			RoleId:    *data.RoleID,
			RoleName:  *data.RoleName,
			LastLogin: lastLogin,
			CreatedAt: timestamppb.New(data.CreatedAt),
			UpdatedAt: timestamppb.New(data.UpdatedAt),
		},
	}, nil
}

func (s *server) GetList(ctx context.Context, request *UserListRequest) (*UserListResponse, error) {
	paging := core.PagingRequest{
		Page:  int(request.GetPage()),
		Limit: int(request.GetLimit()),
	}
	if paging.Invalid() {
		return nil, status.New(codes.InvalidArgument, core.TranslateCtx(ctx, localizations.CommonInvalidData)).Err()
	}

	params := Params{}

	getCount := func() int64 { return s.UserUc.Count(params) }

	getData := func(limit int, offset int) any {
		params.LimitNo = paging.Limit
		params.OffsetNo = offset
		return s.UserUc.GetList(params)
	}
	resp := core.Pagination(paging.Page, paging.Limit, getCount, getData)

	list := []*UserData{}
	for _, u := range resp.List.([]User) {
		var lastLogin *timestamppb.Timestamp
		if u.LastLogin != nil {
			lastLogin = timestamppb.New(*u.LastLogin)
		}

		list = append(list, &UserData{
			Id:        *u.ID,
			Username:  u.Username,
			Email:     u.Email,
			FirstName: u.FirstName,
			LastName:  u.LastName,
			Avatar:    u.Avatar,
			Mobile:    u.Mobile,
			Flag:      int32(u.Flag),
			RoleId:    *u.RoleID,
			RoleName:  *u.RoleName,
			LastLogin: lastLogin,
			CreatedAt: timestamppb.New(u.CreatedAt),
			UpdatedAt: timestamppb.New(u.UpdatedAt),
		})
	}

	return &UserListResponse{
		Code:    code.StatusOK,
		Message: http.StatusText(http.StatusOK),
		Data: &PagingResponse{
			List:  list,
			Page:  int32(resp.Page),
			Limit: int32(resp.Limit),
			Count: int32(resp.Count),
			Total: int32(resp.Total),
			Start: int32(resp.Start),
			End:   int32(resp.End),
		},
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
