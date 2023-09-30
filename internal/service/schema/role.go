package schema

import "github.com/prongbang/user-service/internal/service/database"

type RoleSchema interface {
	Initial()
}

type roleSchema struct {
	DbDriver database.Drivers
}

func (u *roleSchema) Initial() {

}

func NewRoleSchema(dbDriver database.Drivers) RoleSchema {
	return &roleSchema{
		DbDriver: dbDriver,
	}
}
