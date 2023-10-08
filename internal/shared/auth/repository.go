package auth

import "github.com/prongbang/user-service/internal/shared/user"

type Repository interface {
	GetByUsername(username string) user.User
	GetByEmail(email string) user.User
}

type repository struct {
	Ds DataSource
}

func (r *repository) GetByUsername(username string) user.User {
	return r.Ds.GetByUsername(username)
}

func (r *repository) GetByEmail(email string) user.User {
	return r.Ds.GetByEmail(email)
}

func NewRepository(ds DataSource) Repository {
	return &repository{
		Ds: ds,
	}
}
