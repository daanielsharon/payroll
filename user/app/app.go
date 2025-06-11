package app

import (
	"auth/handlers"
	"auth/routes"
	"auth/services"
	"auth/storage"
	"shared/db"

	"github.com/go-chi/chi/v5"
)

func New() *chi.Mux {
	postgres := db.Connect()
	storage := storage.NewStorage(postgres)
	services := services.NewService(storage)
	handlers := handlers.NewHandler(services)

	return routes.InitRoutes(handlers)
}
