package handlers

import (
	"auth/services"
	"encoding/json"
	"net/http"

	httphelper "shared/http"
)

type Handler struct {
	services services.ServiceInterface
}

func NewHandler(services services.ServiceInterface) HandlerInterface {
	return &Handler{services: services}
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var credentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	json.NewDecoder(r.Body).Decode(&credentials)
	token, err := h.services.Login(credentials.Username, credentials.Password)
	if err != nil {
		httphelper.JSONResponse(w, http.StatusBadRequest, "user not found", nil)
		return
	}

	httphelper.JSONResponse(w, http.StatusOK, "Login Successful", token)
}
