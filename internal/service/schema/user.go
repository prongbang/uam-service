package schema

import (
	"context"
	"fmt"
	"github.com/prongbang/uam-service/internal/service/database"
	"github.com/uptrace/bun"
	"time"
)

const tableUsers = "users"

type User struct {
	ID        string    `json:"id" bun:"id,pk,type:uuid,default:uuid_generate_v4()"`
	Username  *string   `json:"username,omitempty"`
	Password  string    `json:"password,omitempty"`
	Email     *string   `json:"email,omitempty"`
	FirstName *string   `json:"first_name,omitempty"`
	LastName  *string   `json:"last_name,omitempty"`
	Avatar    *string   `json:"avatar,omitempty"`
	Mobile    *string   `json:"mobile,omitempty"`
	Flag      int       `json:"flag"`
	LastLogin time.Time `json:"last_login"`
	CreatedAt time.Time `json:"created_at" bun:",default:current_timestamp"`
	UpdatedAt time.Time `json:"updated_at" bun:",default:current_timestamp"`
}

var _ bun.AfterCreateTableHook = (*User)(nil)

func (*User) AfterCreateTable(ctx context.Context, query *bun.CreateTableQuery) error {
	_, err := query.DB().NewCreateIndex().
		Model((*User)(nil)).Table(tableUsers).IfNotExists().
		Index("username_idx").Column("username").
		Index("email_idx").Column("email").
		Exec(ctx)
	return err
}

type UserSchema interface {
	Initial()
}

type userSchema struct {
	DbDriver database.Drivers
}

func (u *userSchema) Initial() {
	ctx := context.Background()
	db := u.DbDriver.GetPqDB()

	_, err := db.NewCreateTable().Model((*User)(nil)).Table(tableUsers).IfNotExists().Exec(ctx)
	if err != nil {
		fmt.Println("Can't create table", tableUsers, err)
	} else {
		fmt.Println(fmt.Sprintf("Table %s created", tableUsers))
	}
}

func NewUserSchema(dbDriver database.Drivers) UserSchema {
	return &userSchema{
		DbDriver: dbDriver,
	}
}
