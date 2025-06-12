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
	"shared/utils"
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
		httphelper.JSONResponse(w, http.StatusBadRequest, "Invalid overtime data", nil)
		return
	}

	overtimeDate, err := time.Parse("2006-01-02", overtimeRequest.Date)
	if err != nil {
		span.AddEvent("Invalid overtime date")
		httphelper.JSONResponse(w, http.StatusBadRequest, "Invalid overtime date", nil)
		return
	}

	filteredOvertime := models.Overtime{
		Date:  overtimeDate,
		Hours: overtimeRequest.Hours,
	}

	err = h.services.Submit(ctx, filteredOvertime)
	if err != nil {
		httphelper.JSONResponse(w, http.StatusInternalServerError, "Overtime run failed", nil)
		return
	}

	httphelper.JSONResponse(w, http.StatusOK, "Overtime run successful", overtimeRequest)
}

func (h *Handler) UpdateOvertimePayroll(w http.ResponseWriter, r *http.Request) {
	tracer := otel.Tracer(fmt.Sprintf("%s/handler", constant.ServiceOvertime))
	ctx, span := tracer.Start(r.Context(), "UpdateOvertimePayroll Handler")
	defer span.End()

	startDate := httphelper.GetQueryParams(r, "startDate")
	endDate := httphelper.GetQueryParams(r, "endDate")
	payrollRunId := httphelper.GetQueryParams(r, "payrollRunId")

	startDateParsed := utils.ConvertStringToDate(startDate)
	endDateParsed := utils.ConvertStringToDate(endDate)
	payrollRunIdParsed, _ := utils.ParseUUID(payrollRunId)

	span.AddEvent("Updating payroll")
	err := h.services.UpdateOvertimePayroll(ctx, startDateParsed, endDateParsed, payrollRunIdParsed)
	if err != nil {
		httphelper.JSONResponse(w, http.StatusInternalServerError, "Overtime payroll update failed", nil)

		log.Println("Overtime payroll update failed", err)
		return
	}

	httphelper.JSONResponse(w, http.StatusOK, "Overtime payroll update successful", nil)
}
