//go:build wireinject
// +build wireinject

package grpc

import (
	"github.com/google/wire"
	"github.com/prongbang/user-service/internal/database"
	"github.com/prongbang/user-service/internal/user"
)

func CreateGRPC(dbDriver database.Drivers) GRPC {
	wire.Build(
		NewGRPC,
		user.ProviderSet,
	)
	return nil
}
