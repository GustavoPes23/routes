package drivers

import "trail/application/domain"

type RoutesService interface {
	GetRouteByID(id string) (domain.Route, error)
	GetAllRoutes() ([]domain.Route, error)
}
