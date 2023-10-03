package schema

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/prongbang/user-service/internal/service/database"
)

const tableUsersRole = "user_roles"

type UserRole struct {
	ID     uuid.UUID `json:"id" bun:"id,pk,type:uuid,default:uuid_generate_v4()"`
	UserID uuid.UUID `json:"user_id" bun:"user_id,type:uuid"`
	RoleID uuid.UUID `json:"role_id" bun:"role_id,type:uuid"`
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