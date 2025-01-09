package repository

import (
	"errors"
	"trail/application/domain"
)

type RouteRepository struct {
	routes map[string]domain.Route
}

func NewRouteRepository() *RouteRepository {
	return &RouteRepository{
		routes: map[string]domain.Route{
			"1": {ID: "1", Name: "Trilha do Lago", Distance: 5.5, Difficulty: "Fácil"},
			"2": {ID: "2", Name: "Caminho do Parque", Distance: 10.0, Difficulty: "Médio"},
			"3": {ID: "3", Name: "Rota das Montanhas", Distance: 20.0, Difficulty: "Difícil"},
		},
	}
}

func (r *RouteRepository) GetAllRoutes() ([]domain.Route, error) {
	var routes []domain.Route
	for _, route := range r.routes {
		routes = append(routes, route)
	}
	return routes, nil
}

func (r *RouteRepository) GetRouteByID(id string) (domain.Route, error) {
	route, exists := r.routes[id]
	if !exists {
		return domain.Route{}, errors.New("rota não encontrada")
	}
	return route, nil
}
