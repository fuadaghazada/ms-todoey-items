package handler

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

type HealthHandler struct {}

func NewHealthHandler(router *chi.Mux) *HealthHandler {
	handler := &HealthHandler{}

	router.Get("/health", handler.Health)
	router.Get("/readiness", handler.Health)

	return handler
}

func (*HealthHandler) Health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}