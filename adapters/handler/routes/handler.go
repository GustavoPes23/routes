package handler

import (
	"net/http"
	"trilha/ports/drivens"

	. "github.com/tbxark/g4vercel"
)

type Routes struct {
	Service drivens.RouteRepository
}

func NewRoutes(service drivens.RouteRepository) *Routes {
	return &Routes{Service: service}
}

func (h *Routes) Handler(c *Context) {
	routes, err := h.Service.GetAllRoutes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, H{"error": "Something went wrong"})
		return
	}

	c.JSON(http.StatusOK, routes)
}
