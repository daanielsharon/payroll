package routes

import (
	"auth/handlers"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func InitRoutes(handler handlers.HandlerInterface) *chi.Mux {
	router := chi.NewRouter()

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "Hello from Auth Service"}`))
	})

	router.Post("/login", handler.Login)

	return router
}
