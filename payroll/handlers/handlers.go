package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"payroll/services"
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

func (h *Handler) CreatePayrollPeriod(w http.ResponseWriter, r *http.Request) {
	tracer := otel.Tracer(fmt.Sprintf("%s/handler", constant.ServicePayroll))
	ctx, span := tracer.Start(r.Context(), "CreatePayrollPeriod Handler")
	defer span.End()

	var payrollPeriod models.PayrollPeriodRequest
	json.NewDecoder(r.Body).Decode(&payrollPeriod)

	filteredPayrollPeriod := models.PayrollPeriod{
		PeriodName: payrollPeriod.PeriodName,
		StartDate:  utils.ConvertStringToDate(payrollPeriod.StartDate),
		EndDate:    utils.ConvertStringToDate(payrollPeriod.EndDate),
	}

	err := h.services.CreatePayrollPeriod(ctx, filteredPayrollPeriod)
	if err != nil {
		httphelper.JSONResponse(w, http.StatusInternalServerError, "Failed to create payroll period", nil)
		return
	}

	httphelper.JSONResponse(w, http.StatusOK, "Payroll period created successfully", nil)
}

func (h *Handler) CreatePayrollRun(w http.ResponseWriter, r *http.Request) {
	tracer := otel.Tracer(fmt.Sprintf("%s/handler", constant.ServicePayroll))
	ctx, span := tracer.Start(r.Context(), "CreatePayrollRun Handler")
	defer span.End()

	var payrollRun models.PayrollRunRequest
	json.NewDecoder(r.Body).Decode(&payrollRun)

	filteredPayrollRun := models.PayrollRun{
		PeriodID: payrollRun.PeriodID,
		RanAt:    time.Now(),
	}

	span.AddEvent("Creating payroll run")
	err := h.services.CreatePayrollRun(ctx, filteredPayrollRun)
	if err != nil {
		httphelper.JSONResponse(w, http.StatusInternalServerError, "Failed to create payroll run", nil)
		return
	}

	httphelper.JSONResponse(w, http.StatusOK, "Payroll run successful", nil)
}
