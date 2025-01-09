package repositories

import (
	repository "trail/adapters/repositories/routes"
	"trail/application/ports/drivens"

	"go.uber.org/fx"
)

var Module = fx.Provide(NewRouteRepository)

func NewRouteRepository() drivens.RoutesRepository {
	return repository.NewRouteRepository()
}
