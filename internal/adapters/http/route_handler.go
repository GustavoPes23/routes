package http

import (
	"encoding/json"
	"net/http"
	"trilha/internal/ports/drivens"

	"github.com/gorilla/mux"
)

type RouteHandler struct {
	Service drivens.RouteRepository
}

func NewRouteHandler(service drivens.RouteRepository) *RouteHandler {
	return &RouteHandler{Service: service}
}

func (h *RouteHandler) GetRoutesHandler(w http.ResponseWriter, r *http.Request) {
	routes, err := h.Service.GetAllRoutes()
	if err != nil {
		http.Error(w, "Erro ao buscar rotas", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(routes)
}

func (h *RouteHandler) GetRouteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	route, err := h.Service.GetRouteByID(id)
	if err != nil {
		http.Error(w, "Rota não encontrada", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(route)
}
