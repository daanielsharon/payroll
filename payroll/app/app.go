package app

import (
	"payroll/handlers"
	"payroll/routes"
	"payroll/services"
	"payroll/storage"
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
