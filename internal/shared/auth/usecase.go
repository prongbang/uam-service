package auth

import (
	"errors"
	"github.com/prongbang/uam-service/internal/localizations"
	"github.com/prongbang/uam-service/internal/pkg/token"
	"github.com/prongbang/uam-service/internal/shared/role"
	"github.com/prongbang/uam-service/internal/shared/user"
	"github.com/prongbang/uam-service/pkg/cryptox"
	"time"
)

type UseCase interface {
	Login(data Login) (Credential, error)
	LoginWithEmail(data Login) (Credential, error)
	LoginWithUsername(data Login) (Credential, error)
}

type useCase struct {
	Repo   Repository
	RoleUc role.UseCase
	UserUc user.UseCase
}

func (u *useCase) Login(data Login) (Credential, error) {
	if data.Email != "" {
		return u.LoginWithEmail(data)
	}
	return u.LoginWithUsername(data)
}

func (u *useCase) LoginWithEmail(data Login) (Credential, error) {
	usr := u.Repo.GetByEmail(data.Email)
	if usr.Password == "" {
		return Credential{}, errors.New(localizations.CommonInvalidData)
	}

	valid := cryptox.VerifyPassword(data.Password, usr.Password)
	if valid {
		_ = u.UserUc.UpdateLastLogin(*usr.ID)

		return u.GetCredentialByUserId(*usr.ID)
	}

	return Credential{}, errors.New(localizations.CommonInvalidData)
}

func (u *useCase) GetCredentialByUserId(userId string) (Credential, error) {
	roles := u.RoleUc.GetByUserIdStringList(userId)
	payload := token.Claims{
		Exp:   time.Now().AddDate(0, 0, 1).Unix(),
		Sub:   userId,
		Roles: roles,
	}
	key, _ := token.GetKeyBytes()
	accessToken, err := token.New(payload, key)
	if err != nil {
		return Credential{}, errors.New(localizations.CommonInvalidData)
	}

	credential := Credential{
		Token: accessToken,
		Roles: roles,
	}
	return credential, nil
}

func (u *useCase) LoginWithUsername(data Login) (Credential, error) {
	usr := u.Repo.GetByUsername(data.Username)
	if usr.Password == "" {
		return Credential{}, errors.New(localizations.CommonInvalidData)
	}

	valid := cryptox.VerifyPassword(data.Password, usr.Password)
	if valid {
		_ = u.UserUc.UpdateLastLogin(*usr.ID)

		return u.GetCredentialByUserId(*usr.ID)
	}

	return Credential{}, errors.New(localizations.CommonInvalidData)
}

func NewUseCase(repo Repository, roleUc role.UseCase, userUc user.UseCase) UseCase {
	return &useCase{
		Repo:   repo,
		RoleUc: roleUc,
		UserUc: userUc,
	}
}
