package user

import (
	"github.com/casbin/casbin/v2"
	"time"
)

type UseCase interface {
	Count(params Params) int64
	GetList(params Params) []User
	GetById(id string) User
	Add(data *User) error
	Update(data *User) error
	UpdatePassword(data *Password) error
	UpdateLastLogin(userId string) error
	Delete(id string) error
}

type useCase struct {
	Repo    Repository
	Enforce *casbin.Enforcer
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

func (u *useCase) Add(data *User) error {
	return u.Repo.Add(data)
}

func (u *useCase) Update(data *User) error {
	return u.Repo.Update(data)
}

func (u *useCase) UpdatePassword(data *Password) error {
	return u.Repo.UpdatePassword(data.UserID, data.NewPassword)
}

func (u *useCase) UpdateLastLogin(userId string) error {
	lastLogin := time.Now()
	return u.Repo.Update(&User{
		ID:        &userId,
		LastLogin: &lastLogin,
	})
}

func (u *useCase) Delete(id string) error {
	return u.Repo.Delete(id)
}

func NewUseCase(repo Repository, enforce *casbin.Enforcer) UseCase {
	return &useCase{
		Repo:    repo,
		Enforce: enforce,
	}
}
