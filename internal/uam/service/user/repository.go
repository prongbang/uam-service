package user

import (
	"github.com/prongbang/uam-service/internal/uam/service/user_creator"
)

type Repository interface {
	Count(params Params) int64
	GetList(params Params) []User
	GetById(params ParamsGetById) User
	GetByEmail(email string) BasicUser
	GetByUsername(username string) BasicUser
	Add(data *CreateUser) error
	Update(data *UpdateUser) error
	UpdatePassword(userId string, password string) error
	Delete(data DeleteUser) error
}

type repository struct {
	Ds            DataSource
	UserCreatorDs user_creator.DataSource
}

func (r *repository) Count(params Params) int64 {
	return r.Ds.Count(params)
}

func (r *repository) GetList(params Params) []User {
	return r.Ds.GetList(params)
}

func (r *repository) GetById(params ParamsGetById) User {
	return r.Ds.GetById(params)
}

func (r *repository) GetByEmail(email string) BasicUser {
	return r.Ds.GetByEmail(email)
}

func (r *repository) GetByUsername(username string) BasicUser {
	return r.Ds.GetByUsername(username)
}

func (r *repository) Add(data *CreateUser) error {
	tx1, err := r.Ds.AddTx(data)
	if err == nil {
		if tx1.Commit() == nil {
			uc := user_creator.CreateUserCreator{UserID: *data.ID, CreatedBy: data.CreatedBy}
			tx2, err2 := r.UserCreatorDs.AddTx(&uc)
			if err2 == nil {
				return tx2.Commit()
			}
			return err2
		}
		return tx1.Rollback()
	}
	return err
}

func (r *repository) Update(data *UpdateUser) error {
	return r.Ds.Update(data)
}

func (r *repository) UpdatePassword(userId string, password string) error {
	return r.Ds.UpdatePassword(userId, password)
}

func (r *repository) Delete(data DeleteUser) error {
	param1 := &user_creator.DeleteUserCreator{UserID: data.ID, CreatedBy: data.Payload.UserID}
	tx1, err := r.UserCreatorDs.DeleteTx(param1)
	if err == nil {
		if tx1.Commit() == nil {
			tx2, err2 := r.Ds.DeleteTx(data.ID)
			if err2 == nil {
				return tx2.Commit()
			}
			return err2
		}
		return tx1.Rollback()
	}
	return err
}

func NewRepository(ds DataSource, userCreatorDs user_creator.DataSource) Repository {
	return &repository{
		Ds:            ds,
		UserCreatorDs: userCreatorDs,
	}
}
