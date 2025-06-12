package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reimbursement/services"
	"shared/constant"
	httphelper "shared/http"
	"shared/models"

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
		httphelper.JSONResponse(w, http.StatusBadRequest, "Invalid reimbursement request", nil)
		return
	}

	filteredReimbursementRequest := models.Reimbursement{
		Amount:      reimbursementRequest.Amount,
		Description: reimbursementRequest.Description,
	}

	err = h.services.Submit(ctx, filteredReimbursementRequest)
	if err != nil {
		httphelper.JSONResponse(w, http.StatusInternalServerError, "Reimbursement run failed", nil)
		return
	}

	httphelper.JSONResponse(w, http.StatusOK, "Reimbursement run successful", reimbursementRequest)
}
