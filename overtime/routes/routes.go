package routes

import (
	"net/http"
	"overtime/handlers"

	"github.com/go-chi/chi/v5"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"

	"shared/constant"
	httphelper "shared/http"
	"shared/router"
)

func InitRoutes(handler handlers.HandlerInterface) *chi.Mux {
	r := router.NewBaseRouter()
	r.Use(otelhttp.NewMiddleware(constant.ServiceOvertime))
	r.Use(httphelper.RequestMiddleware)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		httphelper.JSONResponse(w, http.StatusOK, "Hello from Overtime Service", nil)
	})

	r.Group(func(r chi.Router) {
		r.Use(httphelper.AuthMiddleware)
		r.Use(httphelper.JSONOnly)
		r.With(httphelper.IsEmployee).Post("/submit", handler.Submit)
		r.Post("/update-payroll", handler.UpdateOvertimePayroll)
	})

	return r
}
