package handlers

import (
	"auth/services"
	"encoding/json"
	"fmt"
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

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	tracer := otel.Tracer(fmt.Sprintf("%s/handler", constant.ServiceAuth))
	ctx, span := tracer.Start(r.Context(), "Login Handler")
	defer span.End()

	var credentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	json.NewDecoder(r.Body).Decode(&credentials)
	token, err := h.services.Login(ctx, credentials.Username, credentials.Password)
	if err != nil {
		httphelper.JSONResponse(w, http.StatusBadRequest, "user not found", nil)
		return
	}

	httphelper.JSONResponse(w, http.StatusOK, "Login Successful", token)
}
