package routes

import (
	"auth/handlers"
	"net/http"

	"github.com/go-chi/chi/v5"

	"shared/router"
)

func InitRoutes(handler handlers.HandlerInterface) *chi.Mux {
	r := router.NewBaseRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "Hello from Auth Service"}`))
	})

	r.Post("/login", handler.Login)

	return r
}
