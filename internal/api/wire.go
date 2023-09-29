//go:build wireinject
// +build wireinject

package api

import (
	"github.com/google/wire"
	"github.com/prongbang/user-service/internal/database"
	"github.com/prongbang/user-service/internal/user"
)

func CreateAPI(dbDriver database.Drivers) API {
	wire.Build(
		NewAPI,
		NewRouters,
		user.ProviderSet,
	)
	return nil
}
