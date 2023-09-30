package user

import "github.com/casbin/casbin/v2"

type UseCase interface {
}

type useCase struct {
	Repo    Repository
	Enforce *casbin.Enforcer
}

func NewUseCase(repo Repository, enforce *casbin.Enforcer) UseCase {
	return &useCase{
		Repo:    repo,
		Enforce: enforce,
	}
}
