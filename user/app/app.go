package app

import (
	"shared/db"
	"user/handlers"
	"user/routes"
	"user/services"
	"user/storage"

	"github.com/go-chi/chi/v5"
)

func New() *chi.Mux {
	postgres := db.Connect()
	storage := storage.NewStorage(postgres)
	services := services.NewService(storage)
	handlers := handlers.NewHandler(services)

	return routes.InitRoutes(handlers)
}
