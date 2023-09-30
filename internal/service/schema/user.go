package schema

import "github.com/prongbang/user-service/internal/service/database"

type UserSchema interface {
	Initial()
}

type userSchema struct {
	DbDriver database.Drivers
}

func (u *userSchema) Initial() {

}

func NewUserSchema(dbDriver database.Drivers) UserSchema {
	return &userSchema{
		DbDriver: dbDriver,
	}
}
