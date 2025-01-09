package handlers

import (
	route "trail/adapters/handlers/route"
	routes "trail/adapters/handlers/routes"
	"trail/application/ports/drivers"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"handlers",
	fx.Provide(
		newEngine,
		newRouteHandler,
		newRoutesHandler,
	),
	fx.Invoke(RegisterRoutes),
)

type RoutesDependencies struct {
	fx.In

	RoutesServices drivers.RoutesService
}

func newEngine() *gin.Engine {
	return gin.Default()
}

func newRouteHandler(deps RoutesDependencies) *route.Route {
	return route.NewRoute(deps.RoutesServices)
}

func newRoutesHandler(deps RoutesDependencies) *routes.Routes {
	return routes.NewRoutes(deps.RoutesServices)
}

func RegisterRoutes(server *gin.Engine, routeHandler *route.Route, routesHandler *routes.Routes) {
	server.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello from Go with Uber FX and Gin!",
		})
	})

	server.GET(routesHandler.Pattern(), routesHandler.Handle)
	server.GET(routeHandler.Pattern(), routeHandler.Handle)
}
