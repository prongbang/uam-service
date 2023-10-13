package role

import (
	"github.com/prongbang/uam-service/pkg/core"
	"github.com/uptrace/bun"
)

type Role struct {
	bun.BaseModel `bun:"table:roles,alias:r"`
	ID            string `json:"id" bun:"id,pk,type:uuid"`
	Name          string `json:"name" bun:"name" validate:"required"`
	Level         int    `json:"level" bun:"level"`
}

type CreateRole struct {
	bun.BaseModel `bun:"table:roles,alias:r"`
	ID            *string `json:"id" bun:"id,pk,type:uuid"`
	Name          string  `json:"name" bun:"name" validate:"required"`
	Level         int     `json:"level" bun:"level" validate:"required"`
}

type UpdateRole struct {
	bun.BaseModel `bun:"table:roles,alias:r"`
	ID            string `json:"id" bun:"id,pk,type:uuid"`
	Name          string `json:"name" bun:"name" validate:"required"`
	Level         int    `json:"level" bun:"level"`
}

type BodyIdRequest struct {
	ID string `json:"id"`
}

type GetByIdRequest struct {
	BodyIdRequest
}

type DelByIdRequest struct {
	BodyIdRequest
}

type Filter struct {
	core.Params
}

type Permission struct {
}
