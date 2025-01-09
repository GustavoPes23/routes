package adapters

import (
	"trail/adapters/handlers"
	"trail/adapters/repositories"

	"go.uber.org/fx"
)

var Module = fx.Options(
	handlers.Module,
	repositories.Module,
)
