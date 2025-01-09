package application

import (
	"trail/application/ports/drivens"
	"trail/application/ports/drivers"
	services "trail/application/services/routes"

	"go.uber.org/fx"
)

var Module = fx.Provide(NewRouteService)

func NewRouteService(rep drivens.RoutesRepository) drivers.RoutesService {
	return services.NewRouteService(rep)
}
