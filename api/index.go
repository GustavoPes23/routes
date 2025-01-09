package handler

import (
	"net/http"

	module "trilha/adapters/handler"

	. "github.com/tbxark/g4vercel"
)

var server = New()

func Handler(w http.ResponseWriter, r *http.Request) {
	server.GET("/", func(c *Context) {
		c.JSON(http.StatusOK, H{
			"message": "Hello from Go with g4vercel!",
		})
	})

	server.GET("/routes", module.Routes().Handler)
	server.GET("/routes/:id", module.Route().Handler)
	server.Handle(w, r)
}
