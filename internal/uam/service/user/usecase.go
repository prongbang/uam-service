package user

import (
	"errors"
	"github.com/prongbang/uam-service/internal/localizations"
	"github.com/prongbang/uam-service/internal/uam/service/permissions"
	"github.com/prongbang/uam-service/pkg/code"
	"github.com/prongbang/uam-service/pkg/core"
	"github.com/prongbang/uam-service/pkg/cryptox"
	"time"
)

type UseCase interface {
	Count(params Params) int64
	GetList(params Params) []User
	GetById(params ParamsGetById) User
	Add(data *CreateUser) (User, *core.Error)
	Update(data *UpdateUser) (User, *core.Error)
	UpdatePassword(data *Password) error
	UpdateLastLogin(userId string) error
	Delete(id string) error
}

type useCase struct {
	Repo    Repository
	PermsUc permissions.UseCase
}

func (u *useCase) Count(params Params) int64 {
	if u.PermsUc.IsRoot(params.Payload.Roles, permissions.UamPermissionUsers) {
		params.Permission = permissions.All
	}

	return u.Repo.Count(params)
}

func (u *useCase) GetList(params Params) []User {
	if u.PermsUc.IsRoot(params.Payload.Roles, permissions.UamPermissionUsers) {
		params.Permission = permissions.All
	}

	return u.Repo.GetList(params)
}

func (u *useCase) GetById(params ParamsGetById) User {
	// Check permissions
	if u.PermsUc.IsRoot(params.Payload.Roles, permissions.UamPermissionUsers) {
		params.Permission = permissions.All
	} else if u.PermsUc.Enforces(params.Payload.Roles, permissions.UamPermissionUsers, permissions.Read) {
		// Check my user
		if u.PermsUc.Enforces(params.Payload.Roles, permissions.UamPermissionUsers, permissions.ReadMe) && params.ID != params.Payload.UserID {
			return User{}
		}
	}

	return u.Repo.GetById(params)
}

func (u *useCase) Add(data *CreateUser) (User, *core.Error) {
	// Check permissions
	if !u.PermsUc.IsRoot(data.Payload.Roles, permissions.UamPermissionUsers) {
		if !u.PermsUc.Enforces(data.Payload.Roles, permissions.UamPermissionUsers, permissions.Create) {
			return User{}, &core.Error{Code: code.StatusPermissionDenied, Message: localizations.CommonPermissionDenied}
		}
	}

	if data.Email != "" {
		if rs := u.Repo.GetByEmail(data.Email); core.IsUuid(&rs.ID) {
			return User{}, &core.Error{Code: code.StatusDataDuplicated, Message: localizations.CommonDataDuplicated}
		}
	}
	if data.Username != "" {
		if rs := u.Repo.GetByUsername(data.Username); core.IsUuid(&rs.ID) {
			return User{}, &core.Error{Code: code.StatusDataDuplicated, Message: localizations.CommonDataDuplicated}
		}
	}

	err := u.Repo.Add(data)
	if err != nil {
		return User{}, &core.Error{Code: code.StatusDataInvalid, Message: err.Error()}
	}

	usr := u.GetById(ParamsGetById{ID: *data.ID, Payload: data.Payload})

	return usr, nil
}

func (u *useCase) Update(data *UpdateUser) (User, *core.Error) {
	// Check permissions
	if !u.PermsUc.IsRoot(data.Payload.Roles, permissions.UamPermissionUsers) {
		if u.PermsUc.Enforces(data.Payload.Roles, permissions.UamPermissionUsers, permissions.Update) {
			// Check my user
			if u.PermsUc.Enforces(data.Payload.Roles, permissions.UamPermissionUsers, permissions.UpdateMe) && data.ID != data.Payload.UserID {
				return User{}, &core.Error{Code: code.StatusPermissionDenied, Message: localizations.CommonPermissionDenied}
			}

			// Check user under
			us := u.GetById(ParamsGetById{ID: data.ID, Payload: data.Payload})
			if us.ID == nil {
				return User{}, &core.Error{Code: code.StatusPermissionDenied, Message: localizations.CommonPermissionDenied}
			}
		}
	}

	if data.Email != "" {
		if rs := u.Repo.GetByEmail(data.Email); core.IsUuid(&rs.ID) {
			if rs.ID != data.ID {
				return User{}, &core.Error{Code: code.StatusDataDuplicated, Message: localizations.CommonDataDuplicated}
			}
		}
	}
	if data.Username != "" {
		if rs := u.Repo.GetByUsername(data.Username); core.IsUuid(&rs.ID) {
			if rs.ID != data.ID {
				return User{}, &core.Error{Code: code.StatusDataDuplicated, Message: localizations.CommonDataDuplicated}
			}
		}
	}

	// Reset sensitive data
	data.Password = ""

	err := u.Repo.Update(data)
	if err != nil {
		return User{}, &core.Error{Code: code.StatusDataInvalid, Message: err.Error()}
	}

	usr := u.GetById(ParamsGetById{ID: data.ID, Payload: data.Payload})

	return usr, nil
}

func (u *useCase) UpdatePassword(data *Password) error {
	usr := u.GetById(ParamsGetById{ID: data.UserID})
	if core.IsUuid(usr.ID) && cryptox.VerifyPassword(data.CurrentPassword, usr.Password) {
		return u.Repo.UpdatePassword(data.UserID, data.NewPassword)
	}
	return errors.New(localizations.CommonInvalidData)
}

func (u *useCase) UpdateLastLogin(userId string) error {
	lastLogin := time.Now()
	return u.Repo.Update(&UpdateUser{
		ID:        userId,
		LastLogin: &lastLogin,
	})
}

func (u *useCase) Delete(id string) error {
	return u.Repo.Delete(id)
}

func NewUseCase(repo Repository, permsUc permissions.UseCase) UseCase {
	return &useCase{
		Repo:    repo,
		PermsUc: permsUc,
	}
}
