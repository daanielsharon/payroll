package routes

import (
	"net/http"
	"user/handlers"

	"github.com/go-chi/chi/v5"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"

	"shared/constant"
	"shared/router"
)

func InitRoutes(handler handlers.HandlerInterface) *chi.Mux {
	r := router.NewBaseRouter()
	r.Use(otelhttp.NewMiddleware(constant.ServiceUser))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "Hello from User Service"}`))
	})

	r.Get("/user", handler.GetUserByUsername)

	return r
}
