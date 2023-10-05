package uam

import (
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	NewRouter,
	NewServer,
	NewListener,
)
