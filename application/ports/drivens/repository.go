package drivens

import "trail/application/domain"

type RoutesRepository interface {
	GetAllRoutes() ([]domain.Route, error)
	GetRouteByID(id string) (domain.Route, error)
}
