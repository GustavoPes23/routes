package drivens

import "trilha/domain"

type RouteRepository interface {
	GetAllRoutes() ([]domain.Route, error)
	GetRouteByID(id string) (domain.Route, error)
}
