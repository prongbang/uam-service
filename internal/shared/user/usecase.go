package user

import "github.com/casbin/casbin/v2"

type UseCase interface {
	Count(params Params) int64
	GetList(params Params) []User
	GetById(id string) User
	Add(data *User) error
	Update(data *User) error
	UpdatePassword(data *Password) error
	UpdateLastLogin(userId string) error
	Delete(id string) error
	IsUserUnder(userId1 string, userId2 string) bool
}

type useCase struct {
	Repo    Repository
	Enforce *casbin.Enforcer
}

func (u *useCase) Count(params Params) int64 {
	//TODO implement me
	panic("implement me")
}

func (u *useCase) GetList(params Params) []User {
	//TODO implement me
	panic("implement me")
}

func (u *useCase) GetById(id string) User {
	//TODO implement me
	panic("implement me")
}

func (u *useCase) Add(data *User) error {
	//TODO implement me
	panic("implement me")
}

func (u *useCase) Update(data *User) error {
	//TODO implement me
	panic("implement me")
}

func (u *useCase) UpdatePassword(data *Password) error {
	//TODO implement me
	panic("implement me")
}

func (u *useCase) UpdateLastLogin(userId string) error {
	//TODO implement me
	panic("implement me")
}

func (u *useCase) Delete(id string) error {
	//TODO implement me
	panic("implement me")
}

func (u *useCase) IsUserUnder(userId1 string, userId2 string) bool {
	//TODO implement me
	panic("implement me")
}

func NewUseCase(repo Repository, enforce *casbin.Enforcer) UseCase {
	return &useCase{
		Repo:    repo,
		Enforce: enforce,
	}
}
