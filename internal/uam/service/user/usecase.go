package user

import (
	"errors"
	"github.com/prongbang/uam-service/internal/localizations"
	"github.com/prongbang/uam-service/pkg/code"
	"github.com/prongbang/uam-service/pkg/core"
	"github.com/prongbang/uam-service/pkg/cryptox"
	"time"
)

type UseCase interface {
	Count(params Params) int64
	GetList(params Params) []User
	GetById(id string) User
	Add(data *CreateUser) *core.Error
	Update(data *UpdateUser) *core.Error
	UpdatePassword(data *Password) error
	UpdateLastLogin(userId string) error
	Delete(id string) error
}

type useCase struct {
	Repo Repository
}

func (u *useCase) Count(params Params) int64 {
	return u.Repo.Count(params)
}

func (u *useCase) GetList(params Params) []User {
	return u.Repo.GetList(params)
}

func (u *useCase) GetById(id string) User {
	return u.Repo.GetById(id)
}

func (u *useCase) Add(data *CreateUser) *core.Error {
	if data.Email != "" {
		if rs := u.Repo.GetByEmail(data.Email); core.IsUuid(rs.ID) {
			return &core.Error{Code: code.StatusDataDuplicated, Message: localizations.CommonDataDuplicated}
		}
	}
	if data.Username != "" {
		if rs := u.Repo.GetByUsername(data.Username); core.IsUuid(rs.ID) {
			return &core.Error{Code: code.StatusDataDuplicated, Message: localizations.CommonDataDuplicated}
		}
	}

	err := u.Repo.Add(data)
	if err != nil {
		return &core.Error{Code: code.StatusDataInvalid, Message: err.Error()}
	}

	return nil
}

func (u *useCase) Update(data *UpdateUser) *core.Error {
	if data.Email != "" {
		if rs := u.Repo.GetByEmail(data.Email); core.IsUuid(rs.ID) {
			if *rs.ID != data.ID {
				return &core.Error{Code: code.StatusDataDuplicated, Message: localizations.CommonDataDuplicated}
			}
		}
	}
	if data.Username != "" {
		if rs := u.Repo.GetByUsername(data.Username); core.IsUuid(rs.ID) {
			if *rs.ID != data.ID {
				return &core.Error{Code: code.StatusDataDuplicated, Message: localizations.CommonDataDuplicated}
			}
		}
	}

	// Reset sensitive data
	data.Password = ""

	err := u.Repo.Update(data)
	if err != nil {
		return &core.Error{Code: code.StatusDataInvalid, Message: err.Error()}
	}

	return nil
}

func (u *useCase) UpdatePassword(data *Password) error {
	usr := u.GetById(data.UserID)
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

func NewUseCase(repo Repository) UseCase {
	return &useCase{
		Repo: repo,
	}
}
