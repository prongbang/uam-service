//go:build wireinject
// +build wireinject

package uam

import (
	"github.com/casbin/casbin/v2"
	"github.com/google/wire"
	"github.com/prongbang/uam-service/internal/uam/database"
	"github.com/prongbang/uam-service/internal/uam/service/auth"
	"github.com/prongbang/uam-service/internal/uam/service/forgot"
	"github.com/prongbang/uam-service/internal/uam/service/role"
	"github.com/prongbang/uam-service/internal/uam/service/user"
	"github.com/prongbang/uam-service/internal/uam/service/user_role"
)

func New(dbDriver database.Drivers, enforce *casbin.Enforcer) Services {
	wire.Build(
		NewService,
		NewAPI,
		NewRouters,
		NewListeners,
		NewGRPC,
		role.ProviderSet,
		user.ProviderSet,
		user_role.ProviderSet,
		auth.ProviderSet,
		forgot.ProviderSet,
	)
	return nil
}
