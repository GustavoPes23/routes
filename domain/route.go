package domain

type Route struct {
	ID         string  `json:"id"`
	Name       string  `json:"name"`
	Distance   float64 `json:"distance"` // KM
	Difficulty string  `json:"difficulty"`
}

type RouteRepository interface {
	GetAllRoutes() ([]Route, error)
	GetRouteByID(id string) (Route, error)
}
