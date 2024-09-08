package main

import (
	"fmt"
	"log"
	gohttp "net/http"
	"trilha/internal/adapters/http"
	"trilha/internal/adapters/repository"
	"trilha/internal/services"

	"github.com/gorilla/mux"
)

func main() {
	repo := repository.NewRouteRepository()
	service := services.NewRouteService(repo)

	routeHandler := http.NewRouteHandler(service)

	r := mux.NewRouter()
	r.HandleFunc("/routes", routeHandler.GetRoutesHandler).Methods("GET")
	r.HandleFunc("/routes/{id:[0-9]+}", routeHandler.GetRouteHandler).Methods("GET")

	fmt.Println("Servidor ouvindo na porta 8080")
	log.Fatal(gohttp.ListenAndServe(":8080", r))
}
