package user

import (
	"github.com/prongbang/user-service/pkg/core"
	"time"
)

type User struct {
	ID         string    `json:"id" bun:"id"`
	Username   string    `json:"username" bun:"username"`
	Password   string    `json:"password,omitempty" bun:"password"`
	Email      string    `json:"email" bun:"email"`
	Firstname  string    `json:"firstName" bun:"first_name"`
	Lastname   string    `json:"lastName" bun:"last_name"`
	LastLogin  time.Time `json:"lastLogin" bun:"last_login"`
	Avatar     string    `json:"avatar" bun:"avatar"`
	Mobile     string    `json:"mobile" bun:"mobile"`
	LockDelete bool      `json:"lockDelete" bun:"lock_delete"`
	Flag       int       `json:"flag" bun:"flag"`
	core.Model
}
