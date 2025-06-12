package handlers

import (
	"attendance/services"
	"fmt"
	"log"
	"net/http"
	"shared/constant"
	httphelper "shared/http"

	"go.opentelemetry.io/otel"
)

type Handler struct {
	services services.ServiceInterface
}

func NewHandler(services services.ServiceInterface) HandlerInterface {
	return &Handler{services: services}
}

func (h *Handler) Attendance(w http.ResponseWriter, r *http.Request) {
	tracer := otel.Tracer(fmt.Sprintf("%s/handler", constant.ServiceAttendance))
	ctx, span := tracer.Start(r.Context(), "Attendance Handler")
	defer span.End()

	span.AddEvent("Attending")
	err := h.services.Attend(ctx)
	if err != nil {
		httphelper.JSONResponse(w, http.StatusInternalServerError, "Attendance run failed", nil)

		log.Println("Attendance run failed", err)
		return
	}

	httphelper.JSONResponse(w, http.StatusOK, "Attendance run successful", nil)
}
