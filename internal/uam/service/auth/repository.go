package auth

import (
	"github.com/prongbang/uam-service/internal/uam/service/user"
)

type Repository interface {
	GetByUsername(username string) user.BasicUser
	GetByEmail(email string) user.BasicUser
}

type repository struct {
	Ds     DataSource
	UserDs user.DataSource
}

func (r *repository) GetByUsername(username string) user.BasicUser {
	return r.UserDs.GetByUsername(username)
}

func (r *repository) GetByEmail(email string) user.BasicUser {
	return r.UserDs.GetByEmail(email)
}

func NewRepository(ds DataSource, userDs user.DataSource) Repository {
	return &repository{
		Ds:     ds,
		UserDs: userDs,
	}
}
