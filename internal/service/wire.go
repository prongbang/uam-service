//go:build wireinject
// +build wireinject

package service

import (
	"github.com/casbin/casbin/v2"
	"github.com/google/wire"
	"github.com/prongbang/user-service/internal/service/database"
	"github.com/prongbang/user-service/internal/service/user"
)

func New(dbDriver database.Drivers, enforce *casbin.Enforcer) Service {
	wire.Build(
		NewService,
		NewAPI,
		NewRouters,
		NewGRPC,
		user.ProviderSet,
	)
	return nil
}
