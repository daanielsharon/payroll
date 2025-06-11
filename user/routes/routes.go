package routes

import (
	"auth/handlers"
	"net/http"

	"github.com/go-chi/chi/v5"

	"shared/router"
)

func InitRoutes(handler handlers.HandlerInterface) *chi.Mux {
	router := router.NewBaseRouter()

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "Hello from User Service"}`))
	})

	router.Post("/user", handler.GetUserByUsername)

	return router
}
