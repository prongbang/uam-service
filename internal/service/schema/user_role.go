package schema

import "github.com/prongbang/user-service/internal/service/database"

type UserRoleSchema interface {
	Initial()
}

type userRoleSchema struct {
	DbDriver database.Drivers
}

func (u *userRoleSchema) Initial() {

}

func NewUserRoleSchema(dbDriver database.Drivers) UserRoleSchema {
	return &userRoleSchema{
		DbDriver: dbDriver,
	}
}
