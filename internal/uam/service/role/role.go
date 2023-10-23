package role

import (
	"github.com/prongbang/uam-service/pkg/core"
	"github.com/uptrace/bun"
)

const (
	RoleAnonymous = "anonymous"
)

type Role struct {
	bun.BaseModel `bun:"table:roles,alias:r"`
	ID            string `json:"id" bun:"id,pk,type:uuid"`
	Name          string `json:"name" bun:"name" validate:"required"`
}

type CreateRole struct {
	bun.BaseModel `bun:"table:roles,alias:r"`
	ID            *string `json:"id" bun:"id,pk,type:uuid"`
	Name          string  `json:"name" bun:"name" validate:"required"`
	Level         int32   `json:"level" bun:"level" validate:"required"`
	CreatedBy     string  `json:"createdBy" bun:"created_by"`
}

type UpdateRole struct {
	bun.BaseModel `bun:"table:roles,alias:r"`
	ID            string `json:"id" bun:"id,pk,type:uuid"`
	Name          string `json:"name" bun:"name" validate:"required"`
	Level         int32  `json:"level" bun:"level"`
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

type Params struct {
	core.Params
	UserID string
}

type ParamsGetById struct {
	ID     string
	UserID string
}

type Permission struct {
}
