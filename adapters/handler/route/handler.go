package handler

import (
	"net/http"
	"trilha/ports/drivens"

	. "github.com/tbxark/g4vercel"
)

type Route struct {
	Service drivens.RouteRepository
}

func NewRoute(service drivens.RouteRepository) *Route {
	return &Route{Service: service}
}

func (h *Route) Handler(c *Context) {
	id := c.Param("id")

	route, err := h.Service.GetRouteByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, H{"error": "This route could not be found"})
		return
	}

	c.JSON(http.StatusOK, route)
}
