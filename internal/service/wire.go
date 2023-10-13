//go:build wireinject
// +build wireinject

package service

import (
	"github.com/casbin/casbin/v2"
	"github.com/google/wire"
	"github.com/prongbang/uam-service/internal/service/database"
	"github.com/prongbang/uam-service/internal/service/uam"
	"github.com/prongbang/uam-service/internal/shared/auth"
	"github.com/prongbang/uam-service/internal/shared/role"
	"github.com/prongbang/uam-service/internal/shared/user"
)

func New(dbDriver database.Drivers, enforce *casbin.Enforcer) Service {
	wire.Build(
		NewService,
		NewAPI,
		NewRouters,
		NewGRPC,
		uam.ProviderSet,
		user.ProviderSet,
		role.ProviderSet,
		auth.ProviderSet,
	)
	return nil
}
