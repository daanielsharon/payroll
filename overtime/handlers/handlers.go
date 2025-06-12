package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"overtime/services"
	"shared/constant"
	httphelper "shared/http"
	"shared/models"
	"time"

	"go.opentelemetry.io/otel"
)

type Handler struct {
	services services.ServiceInterface
}

func NewHandler(services services.ServiceInterface) HandlerInterface {
	return &Handler{services: services}
}

func (h *Handler) Submit(w http.ResponseWriter, r *http.Request) {
	tracer := otel.Tracer(fmt.Sprintf("%s/handler", constant.ServiceOvertime))
	ctx, span := tracer.Start(r.Context(), "Submit Handler")
	defer span.End()

	var overtimeRequest models.OvertimeRequest
	err := json.NewDecoder(r.Body).Decode(&overtimeRequest)
	if err != nil {
		span.AddEvent("Invalid overtime data")
		log.Println("Invalid overtime data", err)
		httphelper.JSONResponse(w, http.StatusBadRequest, "Invalid overtime data", nil)
		return
	}

	overtimeDate, err := time.Parse("2006-01-02", overtimeRequest.Date)
	if err != nil {
		span.AddEvent("Invalid overtime date")
		log.Println("Invalid overtime date", err)
		httphelper.JSONResponse(w, http.StatusBadRequest, "Invalid overtime date", nil)
		return
	}

	filteredOvertime := models.Overtime{
		Date:  overtimeDate,
		Hours: overtimeRequest.Hours,
	}

	err = h.services.Submit(ctx, filteredOvertime)
	fmt.Println("err", err)
	if err != nil {
		httphelper.JSONResponse(w, http.StatusInternalServerError, "Overtime run failed", nil)
		return
	}

	httphelper.JSONResponse(w, http.StatusOK, "Overtime run successful", nil)
}
