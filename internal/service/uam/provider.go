package uam

import (
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	NewHandler,
	NewRouter,
	NewValidate,
	NewServer,
	NewListener,
)
