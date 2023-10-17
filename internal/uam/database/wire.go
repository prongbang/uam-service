//go:build wireinject
// +build wireinject

package database

import (
	"github.com/google/wire"
)

func NewDatabaseDriver() Drivers {
	wire.Build(NewPqDriver, NewDrivers)
	return nil
}
