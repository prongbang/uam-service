package schema

import (
	"context"
	"fmt"
	"github.com/prongbang/uam-service/internal/uam/database"
	"github.com/uptrace/bun"
)

const tableUsersCreators = "users_creators"

type UserCreator struct {
	bun.BaseModel `bun:"table:users_creators,alias:u"`
	ID            string `json:"id" bun:"id,pk,type:uuid,default:uuid_generate_v4()"`
	UserID        string `json:"userId" bun:"user_id,type:uuid"`
	CreatedBy     string `json:"createdBy" bun:"created_by,type:uuid"`
}

type UserCreatorSchema interface {
	Initial()
}

type userCreatorSchema struct {
	DbDriver database.Drivers
}

func (u *userCreatorSchema) Initial() {
	ctx := context.Background()
	db := u.DbDriver.GetPqDB()

	_, err := db.NewCreateTable().Model((*UserCreator)(nil)).Table(tableUsersCreators).IfNotExists().Exec(ctx)
	if err != nil {
		fmt.Println("Can't create table", tableUsersCreators, err)
	} else {
		fmt.Println(fmt.Sprintf("Table %s created", tableUsersCreators))
	}
}

func NewUserCreatorSchema(dbDriver database.Drivers) UserCreatorSchema {
	return &userCreatorSchema{
		DbDriver: dbDriver,
	}
}
