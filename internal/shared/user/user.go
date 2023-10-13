package user

import (
	"github.com/prongbang/uam-service/pkg/core"
	"github.com/uptrace/bun"
	"time"
)

type BodyIdRequest struct {
	ID string `json:"id"`
}

type GetByIdRequest struct {
	BodyIdRequest
}

type User struct {
	bun.BaseModel `bun:"table:users,alias:u"`
	ID            *string   `json:"id" bun:"id"`
	Username      string    `json:"username" bun:"username"`
	Password      string    `json:"password,omitempty" bun:"password"`
	Email         string    `json:"email" bun:"email"`
	Firstname     string    `json:"firstName" bun:"first_name"`
	Lastname      string    `json:"lastName" bun:"last_name"`
	LastLogin     time.Time `json:"lastLogin" bun:"last_login"`
	Avatar        string    `json:"avatar" bun:"avatar"`
	Mobile        string    `json:"mobile" bun:"mobile"`
	Flag          int       `json:"flag" bun:"flag"`
	core.Model
}

type Password struct {
	UserID          int64  `json:"id"`
	CurrentPassword string `json:"currentPassword" validate:"required"`
	NewPassword     string `json:"newPassword" validate:"required"`
}

type Params struct {
	core.Params
}
