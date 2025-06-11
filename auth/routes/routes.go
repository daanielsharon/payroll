package routes

import (
	"auth/handlers"
	"net/http"

	"github.com/go-chi/chi/v5"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"

	"shared/constant"
	"shared/router"
)

func InitRoutes(handler handlers.HandlerInterface) *chi.Mux {
	r := router.NewBaseRouter()
	r.Use(otelhttp.NewMiddleware(constant.ServiceAuth))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "Hello from Auth Service"}`))
	})

	r.Post("/login", handler.Login)

	return r
}
