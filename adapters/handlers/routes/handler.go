package handler

import (
	"net/http"
	"trail/application/ports/drivers"

	"github.com/gin-gonic/gin"
)

type Routes struct {
	Service drivers.RoutesService
}

func NewRoutes(service drivers.RoutesService) *Routes {
	return &Routes{Service: service}
}

func (r *Routes) Pattern() string {
	return "/routes"
}

func (r *Routes) Handle(c *gin.Context) {
	routes, err := r.Service.GetAllRoutes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
		return
	}

	c.JSON(http.StatusOK, routes)
}
