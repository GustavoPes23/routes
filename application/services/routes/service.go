package services

import (
	"trail/application/domain"
	"trail/application/ports/drivens"
)

type RouteService struct {
	repository drivens.RoutesRepository
}

func NewRouteService(repo drivens.RoutesRepository) *RouteService {
	return &RouteService{repository: repo}
}

func (s *RouteService) GetAllRoutes() ([]domain.Route, error) {
	return s.repository.GetAllRoutes()
}

func (s *RouteService) GetRouteByID(id string) (domain.Route, error) {
	return s.repository.GetRouteByID(id)
}
