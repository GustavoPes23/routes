package api

import (
	stdhttp "net/http"
	"trilha/internal/adapters/http"
	"trilha/internal/adapters/repository"
	"trilha/internal/services"

	"github.com/gorilla/mux"
)

var router *mux.Router

func init() {
	// Inicialize dependências
	repo := repository.NewRouteRepository()
	service := services.NewRouteService(repo)
	routeHandler := http.NewRouteHandler(service)

	// Configure o roteador
	router = mux.NewRouter()
	router.HandleFunc("/routes", routeHandler.GetRoutesHandler).Methods("GET")
	router.HandleFunc("/routes/{id:[0-9]+}", routeHandler.GetRouteHandler).Methods("GET")
}

// Handler para a Vercel
func Handler(w stdhttp.ResponseWriter, r *stdhttp.Request) {
	router.ServeHTTP(w, r)
}
