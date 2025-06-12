package app

import (
	"attendance/handlers"
	"attendance/routes"
	"attendance/services"
	"attendance/storage"
	"shared/db"
	"shared/utils"

	"github.com/go-chi/chi/v5"
)

func New() *chi.Mux {
	postgres := db.Connect()
	storage := storage.NewStorage(postgres)
	services := services.NewService(storage, utils.GetCurrentTime)
	handlers := handlers.NewHandler(services)

	return routes.InitRoutes(handlers)
}
