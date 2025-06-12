package routes

import (
	"attendance/handlers"
	"net/http"

	"github.com/go-chi/chi/v5"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"

	"shared/constant"
	httphelper "shared/http"
	"shared/router"
)

func InitRoutes(handler handlers.HandlerInterface) *chi.Mux {
	r := router.NewBaseRouter()
	r.Use(otelhttp.NewMiddleware(constant.ServiceAttendance))
	r.Use(httphelper.RequestMiddleware)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		httphelper.JSONResponse(w, http.StatusOK, "Hello from Attendance Service", nil)
	})

	r.Group(func(r chi.Router) {
		r.Use(httphelper.AuthMiddleware)
		r.Get("/attendance", handler.GetAttendanceByUserIdAndDate)
	})

	r.Group(func(r chi.Router) {
		r.Use(httphelper.AuthMiddleware)
		r.Use(httphelper.JSONOnly)
		r.With(httphelper.IsEmployee).Post("/submit", handler.Submit)
		r.With(httphelper.IsAdmin).Post("/update-payroll", handler.UpdateAttendancePayroll)
	})

	return r
}
