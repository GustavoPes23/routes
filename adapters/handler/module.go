package handler

import (
	route "trilha/adapters/handler/route"
	routes "trilha/adapters/handler/routes"
	repository "trilha/adapters/repository/routes"
	services "trilha/services/routes"
)

var (
	repo    = repository.NewRouteRepository()
	service = services.NewRouteService(repo)
)

func Routes() *routes.Routes {
	return routes.NewRoutes(service)
}

func Route() *route.Route {
	return route.NewRoute(service)
}
