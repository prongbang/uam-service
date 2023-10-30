package auth

import (
	"errors"
	"github.com/prongbang/uam-service/internal/localizations"
	"github.com/prongbang/uam-service/internal/pkg/token"
	"github.com/prongbang/uam-service/internal/uam/service/permissions"
	"github.com/prongbang/uam-service/internal/uam/service/role"
	"github.com/prongbang/uam-service/internal/uam/service/user"
	"github.com/prongbang/uam-service/pkg/cryptox"
	"time"
)

type UseCase interface {
	Login(data Login) (Credential, error)
	LoginWithEmail(data Login) (Credential, error)
	LoginWithUsername(data Login) (Credential, error)
	VerifyToken(accessToken string) error
	RestEnforce(requests ...any) bool
	RbacEnforce(requests ...any) bool
}

type useCase struct {
	Repo    Repository
	RoleUc  role.UseCase
	UserUc  user.UseCase
	PermsUc permissions.UseCase
}

func (u *useCase) RestEnforce(requests ...any) bool {
	allowed, err := u.PermsUc.RestEnforce(requests...)
	if err != nil {
		return false
	}
	return allowed
}

func (u *useCase) RbacEnforce(requests ...any) bool {
	allowed, err := u.PermsUc.RbacEnforce(requests...)
	if err != nil {
		return false
	}
	return allowed
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
		_ = u.UserUc.UpdateLastLogin(usr.ID)

		return u.GetCredentialByUserId(usr.ID)
	}

	return Credential{}, errors.New(localizations.CommonInvalidData)
}

func (u *useCase) GetCredentialByUserId(userId string) (Credential, error) {
	roles := u.RoleUc.GetListByUserIdString(userId)
	payload := token.Claims{
		Exp:    time.Now().AddDate(0, 0, 1).Unix(),
		UserID: userId,
		Roles:  roles,
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
		_ = u.UserUc.UpdateLastLogin(usr.ID)

		return u.GetCredentialByUserId(usr.ID)
	}

	return Credential{}, errors.New(localizations.CommonInvalidData)
}

func (u *useCase) VerifyToken(accessToken string) error {
	_, err := token.Verification(accessToken)
	return err
}

func NewUseCase(repo Repository, roleUc role.UseCase, userUc user.UseCase, permsUc permissions.UseCase) UseCase {
	return &useCase{
		Repo:    repo,
		RoleUc:  roleUc,
		UserUc:  userUc,
		PermsUc: permsUc,
	}
}
