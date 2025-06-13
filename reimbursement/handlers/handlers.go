package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"reimbursement/services"
	"shared/constant"
	httphelper "shared/http"
	"shared/models"
	"shared/utils"

	"go.opentelemetry.io/otel"
)

type Handler struct {
	services services.ServiceInterface
}

func NewHandler(services services.ServiceInterface) HandlerInterface {
	return &Handler{services: services}
}

func (h *Handler) Submit(w http.ResponseWriter, r *http.Request) {
	tracer := otel.Tracer(fmt.Sprintf("%s/handler", constant.ServiceReimbursement))
	ctx, span := tracer.Start(r.Context(), "Submit Handler")
	defer span.End()

	var reimbursementRequest models.ReimbursementRequest
	err := json.NewDecoder(r.Body).Decode(&reimbursementRequest)
	if err != nil {
		span.RecordError(fmt.Errorf("invalid reimbursement request: %v", err))
		httphelper.JSONResponse(w, http.StatusBadRequest, "Invalid reimbursement request", nil)
		return
	}

	date := utils.ConvertStringToDate(reimbursementRequest.Date)
	filteredReimbursementRequest := models.Reimbursement{
		Amount:      reimbursementRequest.Amount,
		Description: reimbursementRequest.Description,
		Date:        date,
	}

	err = h.services.Submit(ctx, filteredReimbursementRequest)
	if err != nil {
		span.RecordError(fmt.Errorf("reimbursement run failed: %v", err))
		httphelper.JSONResponse(w, http.StatusInternalServerError, "Reimbursement run failed", nil)
		return
	}

	httphelper.JSONResponse(w, http.StatusOK, "Reimbursement run successful", reimbursementRequest)
}

func (h *Handler) UpdateReimbursementPayroll(w http.ResponseWriter, r *http.Request) {
	tracer := otel.Tracer(fmt.Sprintf("%s/handler", constant.ServiceReimbursement))
	ctx, span := tracer.Start(r.Context(), "UpdateReimbursementPayroll Handler")
	defer span.End()

	startDate := httphelper.GetQueryParams(r, "startDate")
	endDate := httphelper.GetQueryParams(r, "endDate")
	payrollRunId := httphelper.GetQueryParams(r, "payrollRunId")

	startDateParsed := utils.ConvertStringToDate(startDate)
	endDateParsed := utils.ConvertStringToDate(endDate)
	payrollRunIdParsed, _ := utils.ParseUUID(payrollRunId)

	span.AddEvent("Updating payroll")
	err := h.services.UpdateReimbursementPayroll(ctx, startDateParsed, endDateParsed, payrollRunIdParsed)
	if err != nil {
		httphelper.JSONResponse(w, http.StatusInternalServerError, "Reimbursement payroll update failed", nil)

		log.Println("Reimbursement payroll update failed", err)
		return
	}

	httphelper.JSONResponse(w, http.StatusOK, "Reimbursement payroll update successful", nil)
}
