package handler

import (
	"net/http"
	"trail/application/ports/drivers"

	"github.com/gin-gonic/gin"
)

type Route struct {
	Service drivers.RoutesService
}

func NewRoute(service drivers.RoutesService) *Route {
	return &Route{Service: service}
}

func (r *Route) Pattern() string {
	return "/routes/:id"
}

func (r *Route) Handle(c *gin.Context) {
	id := c.Param("id")

	route, err := r.Service.GetRouteByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Route not found"})
		return
	}

	c.JSON(http.StatusOK, route)
}
