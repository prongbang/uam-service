package schema

import (
	"context"
	"fmt"
	"github.com/prongbang/uam-service/internal/uam/database"
	"github.com/uptrace/bun"
)

const tableUsersRole = "users_roles"

type UserRole struct {
	bun.BaseModel `bun:"table:users_roles,alias:u"`
	ID            string `json:"id" bun:"id,pk,type:uuid,default:uuid_generate_v4()"`
	UserID        string `json:"user_id" bun:"user_id,type:uuid"`
	RoleID        string `json:"role_id" bun:"role_id,type:uuid"`
	CreatedBy     string `json:"created_by" bun:"created_by,type:uuid"`
}

type UserRoleSchema interface {
	Initial()
}

type userRoleSchema struct {
	DbDriver database.Drivers
}

func (u *userRoleSchema) Initial() {
	ctx := context.Background()
	db := u.DbDriver.GetPqDB()

	_, err := db.NewCreateTable().
		Model((*UserRole)(nil)).Table(tableUsersRole).IfNotExists().
		ForeignKey(`("user_id") REFERENCES "` + tableUsers + `" ("id") ON DELETE CASCADE`).
		ForeignKey(`("role_id") REFERENCES "` + tableRoles + `" ("id") ON DELETE CASCADE`).
		Exec(ctx)
	if err != nil {
		fmt.Println("Can't create table", tableUsersRole, err)
	} else {
		fmt.Println(fmt.Sprintf("Table %s created", tableUsersRole))
	}
}

func NewUserRoleSchema(dbDriver database.Drivers) UserRoleSchema {
	return &userRoleSchema{
		DbDriver: dbDriver,
	}
}
