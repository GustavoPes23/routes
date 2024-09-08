package drivens

import "trilha/internal/domain"

type RouteRepository interface {
	GetAllRoutes() ([]domain.Route, error)
	GetRouteByID(id string) (domain.Route, error)
}
