package handlers

import (
	"net/http"
	"reimbursement/services"
	httphelper "shared/http"
)

type Handler struct {
	services services.ServiceInterface
}

func NewHandler(services services.ServiceInterface) HandlerInterface {
	return &Handler{services: services}
}

func (h *Handler) Run(w http.ResponseWriter, r *http.Request) {
	httphelper.JSONResponse(w, http.StatusOK, "Reimbursement run successful", nil)
}
