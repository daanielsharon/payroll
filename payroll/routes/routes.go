package routes

import (
	"net/http"
	"payroll/handlers"

	"github.com/go-chi/chi/v5"

	httphelper "shared/http"
	"shared/router"
)

func InitRoutes(handler handlers.HandlerInterface) *chi.Mux {
	r := router.NewBaseRouter()
	r.Use(httphelper.AuthMiddleware)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "Hello from Payroll Service"}`))
	})

	r.With(httphelper.IsAdmin).Post("/run", handler.Run)

	return r
}
