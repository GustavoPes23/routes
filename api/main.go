package main

import (
	"fmt"
	"log"
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

func main() {
	repo := repository.NewRouteRepository()
	service := services.NewRouteService(repo)

	routeHandler := http.NewRouteHandler(service)

	r := mux.NewRouter()
	r.HandleFunc("/routes", routeHandler.GetRoutesHandler).Methods("GET")
	r.HandleFunc("/routes/{id:[0-9]+}", routeHandler.GetRouteHandler).Methods("GET")

	fmt.Println("Servidor ouvindo na porta 8080")
	log.Fatal(stdhttp.ListenAndServe(":8080", r))
}
