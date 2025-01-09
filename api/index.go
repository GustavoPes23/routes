package handler

import (
	"context"
	"net/http"
	"trail/adapters"
	"trail/application"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

var ginEngine *gin.Engine

func init() {
	ctx := context.Background()

	app := fx.New(
		adapters.Module,
		application.Module,
		fx.Populate(&ginEngine),
		fx.Invoke(startServer),
	)

	if err := app.Start(ctx); err != nil {
		panic(err)
	}
}

func startServer(lifecycle fx.Lifecycle, server *gin.Engine) {
	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				server.Run()
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return nil
		},
	})
}

func Handler(w http.ResponseWriter, r *http.Request) {
	ginEngine.ServeHTTP(w, r)
}
