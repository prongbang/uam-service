package permissions

import "github.com/google/wire"

var ProviderSet = wire.NewSet(
	NewUseCase,
)
