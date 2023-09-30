//go:build wireinject
// +build wireinject

package schema

import (
	"github.com/google/wire"
	"github.com/prongbang/user-service/internal/service/database"
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
