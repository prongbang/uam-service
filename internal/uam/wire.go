//go:build wireinject
// +build wireinject

package uam

import (
	"github.com/google/wire"
	"github.com/prongbang/uam-service/internal/pkg/casbinx"
	"github.com/prongbang/uam-service/internal/uam/database"
	"github.com/prongbang/uam-service/internal/uam/interceptor"
	"github.com/prongbang/uam-service/internal/uam/service/auth"
	"github.com/prongbang/uam-service/internal/uam/service/forgot"
	"github.com/prongbang/uam-service/internal/uam/service/role"
	"github.com/prongbang/uam-service/internal/uam/service/user"
	"github.com/prongbang/uam-service/internal/uam/service/user_creator"
	"github.com/prongbang/uam-service/internal/uam/service/user_role"
)

func New(dbDriver database.Drivers, casbinXs casbinx.CasbinXs) Services {
	wire.Build(
		NewService,
		NewAPI,
		NewRouters,
		NewListeners,
		NewGRPC,
		interceptor.ProviderSet,
		role.ProviderSet,
		user.ProviderSet,
		user_role.ProviderSet,
		user_creator.ProviderSet,
		auth.ProviderSet,
		forgot.ProviderSet,
	)
	return nil
}
