package app

import (
	"auth/handlers"
	"auth/routes"
	"auth/services"

	"github.com/go-chi/chi/v5"
)

func New() *chi.Mux {
	services := services.NewService()
	handlers := handlers.NewHandler(services)

	return routes.InitRoutes(handlers)
}
