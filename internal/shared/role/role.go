package role

import "github.com/prongbang/user-service/pkg/core"

type Role struct {
	ID    string `json:"id" db:"id"`
	Name  string `json:"name" db:"name" validate:"required"`
	Level int    `json:"level" db:"level"`
}

type Filter struct {
	core.Filter
}

type Permission struct {
}
