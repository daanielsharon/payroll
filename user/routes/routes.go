package routes

import (
	"net/http"
	"user/handlers"

	"github.com/go-chi/chi/v5"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"

	"shared/constant"
	httphelper "shared/http"
	"shared/router"
)

func InitRoutes(handler handlers.HandlerInterface) *chi.Mux {
	r := router.NewBaseRouter()
	r.Use(otelhttp.NewMiddleware(constant.ServiceUser))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		httphelper.JSONResponse(w, http.StatusOK, "Hello from User Service", nil)
	})

	r.Group(func(r chi.Router) {
		r.Use(httphelper.JSONOnly)
		r.Get("/user", handler.GetUserByUsername)
	})

	return r
}
