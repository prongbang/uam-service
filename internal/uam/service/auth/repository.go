package auth

import (
	"github.com/prongbang/uam-service/internal/uam/service/user"
)

type Repository interface {
	GetByUsername(username string) user.SensitiveUser
	GetByEmail(email string) user.SensitiveUser
}

type repository struct {
	Ds     DataSource
	UserDs user.DataSource
}

func (r *repository) GetByUsername(username string) user.SensitiveUser {
	return r.UserDs.GetByUsername(username)
}

func (r *repository) GetByEmail(email string) user.SensitiveUser {
	return r.UserDs.GetByEmail(email)
}

func NewRepository(ds DataSource, userDs user.DataSource) Repository {
	return &repository{
		Ds:     ds,
		UserDs: userDs,
	}
}
