package user

import "github.com/prongbang/user-service/pkg/cryptox"

type Repository interface {
	FindByUsername(username, password string) User
	FindByEmail(email, password string) User
}

type repository struct {
	Ds DataSource
}

func (r *repository) FindByUsername(username, password string) User {
	u := r.Ds.GetByUsername(username)
	if cryptox.VerifyPassword(u.Password, password) {

	}
	//TODO implement me
	panic("implement me")
}

func (r *repository) FindByEmail(email, password string) User {
	//TODO implement me
	panic("implement me")
}

func NewRepository(ds DataSource) Repository {
	return &repository{
		Ds: ds,
	}
}
