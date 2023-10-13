package schema

import (
	"context"
	"fmt"
	"github.com/prongbang/uam-service/internal/service/database"
)

const tableRoles = "roles"

type Role struct {
	ID    string `json:"id" bun:"id,pk,type:uuid,default:uuid_generate_v4()"`
	Name  string `json:"name"`
	Level int    `json:"level"`
}

type RoleSchema interface {
	Initial()
}

type roleSchema struct {
	DbDriver database.Drivers
}

func (u *roleSchema) Initial() {
	ctx := context.Background()
	db := u.DbDriver.GetPqDB()

	_, err := db.NewCreateTable().Model((*Role)(nil)).Table(tableRoles).IfNotExists().Exec(ctx)
	if err != nil {
		fmt.Println("Can't create table", tableRoles, err)
	} else {
		fmt.Println(fmt.Sprintf("Table %s created", tableRoles))
	}
}

func NewRoleSchema(dbDriver database.Drivers) RoleSchema {
	return &roleSchema{
		DbDriver: dbDriver,
	}
}
