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

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		httphelper.JSONResponse(w, http.StatusOK, "Hello from Payroll Service", nil)
	})

	r.Group(func(r chi.Router) {
		r.Use(httphelper.AuthMiddleware)
		r.With(httphelper.IsAdmin).Post("/run", handler.Run)
	})

	return r
}
