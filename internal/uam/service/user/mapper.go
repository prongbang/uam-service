package user

import (
	"github.com/prongbang/uam-service/pkg/core"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToUserDataMapper(u User) *UserData {
	var lastLogin *timestamppb.Timestamp
	if u.LastLogin != nil {
		lastLogin = timestamppb.New(*u.LastLogin)
	}

	roles := []*UserRoleResponse{}
	for _, r := range u.Roles {
		roles = append(roles, &UserRoleResponse{Id: r.ID, Name: r.Name})
	}

	return &UserData{
		Id:        *u.ID,
		Username:  u.Username,
		Email:     u.Email,
		FirstName: u.FirstName,
		LastName:  u.FirstName,
		LastLogin: lastLogin,
		Avatar:    u.Avatar,
		Mobile:    u.Mobile,
		Roles:     roles,
		CreatedAt: timestamppb.New(u.CreatedAt),
		UpdatedAt: timestamppb.New(u.UpdatedAt),
	}
}

func ToUserPagingMapper(resp core.Paging[User]) *PagingResponse {
	list := []*UserData{}
	for _, u := range resp.List {
		var lastLogin *timestamppb.Timestamp
		if u.LastLogin != nil {
			lastLogin = timestamppb.New(*u.LastLogin)
		}

		roles := []*UserRoleResponse{}
		for _, r := range u.Roles {
			roles = append(roles, &UserRoleResponse{Id: r.ID, Name: r.Name})
		}

		list = append(list, &UserData{
			Id:        *u.ID,
			Username:  u.Username,
			Email:     u.Email,
			FirstName: u.FirstName,
			LastName:  u.LastName,
			Avatar:    u.Avatar,
			Mobile:    u.Mobile,
			LastLogin: lastLogin,
			Roles:     roles,
			CreatedAt: timestamppb.New(u.CreatedAt),
			UpdatedAt: timestamppb.New(u.UpdatedAt),
		})
	}

	return &PagingResponse{
		List:  list,
		Page:  int32(resp.Page),
		Limit: int32(resp.Limit),
		Count: int32(resp.Count),
		Total: int32(resp.Total),
		Start: int32(resp.Start),
		End:   int32(resp.End),
	}
}
