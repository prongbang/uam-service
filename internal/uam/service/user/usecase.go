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
	UpdatePassword(data Password) error
	UpdatePasswordMe(data MyPassword) error
	UpdateLastLogin(userId string) error
	Delete(data DeleteUser) error
}

type useCase struct {
	Repo    Repository
	PermsUc permissions.UseCase
}

func (u *useCase) Count(params Params) int64 {
	if u.PermsUc.RbacIsRoot(params.Payload.Roles, permissions.UamPermissionUsers) {
		params.Permission = permissions.All
	}

	return u.Repo.Count(params)
}

func (u *useCase) GetList(params Params) []User {
	if u.PermsUc.RbacIsRoot(params.Payload.Roles, permissions.UamPermissionUsers) {
		params.Permission = permissions.All
	}

	return u.Repo.GetList(params)
}

func (u *useCase) GetById(params ParamsGetById) User {
	// Check permissions
	if u.PermsUc.RbacIsRoot(params.Payload.Roles, permissions.UamPermissionUsers) {
		params.Permission = permissions.All
	} else if u.PermsUc.RbacEnforces(params.Payload.Roles, permissions.UamPermissionUsers, permissions.Read) {
		// Check my user
		if u.PermsUc.RbacEnforces(params.Payload.Roles, permissions.UamPermissionUsers, permissions.ReadMe) && params.ID != params.Payload.UserID {
			return User{}
		}
	}

	return u.Repo.GetById(params)
}

func (u *useCase) Add(data *CreateUser) (User, *core.Error) {
	// Check permissions
	if !u.PermsUc.RbacIsRoot(data.Payload.Roles, permissions.UamPermissionUsers) {
		if !u.PermsUc.RbacEnforces(data.Payload.Roles, permissions.UamPermissionUsers, permissions.Create) {
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
	if !u.PermsUc.RbacIsRoot(data.Payload.Roles, permissions.UamPermissionUsers) {
		if u.PermsUc.RbacEnforces(data.Payload.Roles, permissions.UamPermissionUsers, permissions.Update) {
			// Check my user
			if u.PermsUc.RbacEnforces(data.Payload.Roles, permissions.UamPermissionUsers, permissions.UpdateMe) && data.ID != data.Payload.UserID {
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

func (u *useCase) UpdatePassword(data Password) error {
	// Check permissions
	if !u.PermsUc.RbacIsRoot(data.Payload.Roles, permissions.UamPermissionUsers) {
		if u.PermsUc.RbacEnforces(data.Payload.Roles, permissions.UamPermissionUsers, permissions.Update) {
			// Check my user
			if u.PermsUc.RbacEnforces(data.Payload.Roles, permissions.UamPermissionUsers, permissions.UpdateMe) && data.UserID != data.Payload.UserID {
				return errors.New(localizations.CommonPermissionDenied)
			}

			// Check user under
			us := u.GetById(ParamsGetById{ID: data.UserID, Payload: data.Payload})
			if us.ID == nil {
				return errors.New(localizations.CommonPermissionDenied)
			}
		}
	}

	usr := u.Repo.GetSensitiveById(data.UserID)
	if core.IsUuid(&usr.ID) && cryptox.VerifyPassword(data.CurrentPassword, usr.Password) {
		err := u.Repo.UpdatePassword(data.UserID, data.NewPassword)
		return err
	}
	return errors.New(localizations.CommonInvalidData)
}

func (u *useCase) UpdatePasswordMe(data MyPassword) error {
	usr := u.Repo.GetSensitiveById(data.UserID)
	if core.IsUuid(&usr.ID) && cryptox.VerifyPassword(data.CurrentPassword, usr.Password) {
		err := u.Repo.UpdatePassword(data.UserID, data.NewPassword)
		return err
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

func (u *useCase) Delete(data DeleteUser) error {
	// Check permissions
	if !u.PermsUc.RbacIsRoot(data.Payload.Roles, permissions.UamPermissionUsers) {
		if u.PermsUc.RbacEnforces(data.Payload.Roles, permissions.UamPermissionUsers, permissions.Delete) {
			// Check my user
			if data.ID == data.Payload.UserID {
				return errors.New(localizations.CommonPermissionDenied)
			}

			// Check user under
			us := u.GetById(ParamsGetById{ID: data.ID, Payload: data.Payload})
			if us.ID == nil {
				return errors.New(localizations.CommonNotFoundData)
			}
		}
	}

	return u.Repo.Delete(data)
}

func NewUseCase(repo Repository, permsUc permissions.UseCase) UseCase {
	return &useCase{
		Repo:    repo,
		PermsUc: permsUc,
	}
}
