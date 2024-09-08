package services

import "trilha/internal/domain"

type RouteService struct {
	Repo domain.RouteRepository
}

func NewRouteService(repo domain.RouteRepository) *RouteService {
	return &RouteService{Repo: repo}
}

func (s *RouteService) GetAllRoutes() ([]domain.Route, error) {
	return s.Repo.GetAllRoutes()
}

func (s *RouteService) GetRouteByID(id string) (domain.Route, error) {
	return s.Repo.GetRouteByID(id)
}
