//go:build wireinject
// +build wireinject

package schema

import (
	"github.com/google/wire"
	"github.com/prongbang/uam-service/internal/uam/database"
)

func New(dbDriver database.Drivers) Schema {
	wire.Build(
		NewSchema,
		NewUserSchema,
		NewRoleSchema,
		NewUserRoleSchema,
		NewRBACSchema,
	)
	return nil
}
