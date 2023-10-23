package user

import (
	"github.com/prongbang/uam-service/internal/uam/service/role"
	"github.com/prongbang/uam-service/pkg/core"
	"github.com/uptrace/bun"
	"time"
)

const (
	PasswordMin = 8
	UsernameMin = 4
)

type BodyIdRequest struct {
	ID string `json:"id"`
}

type GetByIdRequest struct {
	BodyIdRequest
}

type DeleteByIdRequest struct {
	BodyIdRequest
}

type User struct {
	bun.BaseModel `bun:"table:users,alias:u"`
	ID            *string     `json:"id" bun:"id,pk,type:uuid"`
	Username      string      `json:"username" bun:"username"`
	Password      string      `json:"password,omitempty" bun:"password"`
	Email         string      `json:"email" bun:"email"`
	FirstName     string      `json:"firstName" bun:"first_name"`
	LastName      string      `json:"lastName" bun:"last_name"`
	LastLogin     *time.Time  `json:"lastLogin" bun:"last_login"`
	Avatar        string      `json:"avatar" bun:"avatar"`
	Mobile        string      `json:"mobile" bun:"mobile"`
	Flag          int         `json:"-" bun:"flag"`
	RolesJson     string      `json:"-" bun:"roles_json"`
	Roles         []role.Role `json:"roles" bun:"-"`
	core.Model
}

type CreateUser struct {
	bun.BaseModel `bun:"table:users,alias:u"`
	ID            *string `json:"id" bun:"id,pk,type:uuid"`
	Username      string  `json:"username" bun:"username"`
	Password      string  `json:"password" bun:"password" validate:"required"`
	Email         string  `json:"email" bun:"email"`
	FirstName     string  `json:"firstName" bun:"first_name"`
	LastName      string  `json:"lastName" bun:"last_name"`
	Avatar        string  `json:"avatar" bun:"avatar"`
	Mobile        string  `json:"mobile" bun:"mobile"`
	CreatedBy     string  `json:"createdBy" bun:"-"`
}

type UpdateUser struct {
	bun.BaseModel `bun:"table:users,alias:u"`
	ID            string     `json:"id" bun:"id,pk,type:uuid"`
	Username      string     `json:"username" bun:"username"`
	Password      string     `json:"-"`
	Email         string     `json:"email" bun:"email"`
	FirstName     string     `json:"firstName" bun:"first_name"`
	LastName      string     `json:"lastName" bun:"last_name"`
	Avatar        string     `json:"avatar" bun:"avatar"`
	Mobile        string     `json:"mobile" bun:"mobile"`
	LastLogin     *time.Time `json:"-" bun:"last_login"`
	Flag          int        `json:"-" bun:"flag"`
}

type Password struct {
	UserID          string `json:"userId" validate:"required"`
	NewPassword     string `json:"newPassword" validate:"required"`
	CurrentPassword string `json:"currentPassword" validate:"required"`
}

type MyPassword struct {
	UserID          string `json:"userId"`
	NewPassword     string `json:"newPassword" validate:"required"`
	CurrentPassword string `json:"currentPassword" validate:"required"`
}

type Params struct {
	core.Params
	UserID string
}

type ParamsGetById struct {
	ID     string
	UserID string
}
