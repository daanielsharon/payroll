package handlers

import (
	"auth/services"
	"net/http"

	httphelper "shared/http"
)

type Handler struct {
	services services.ServiceInterface
}

func NewHandler(services services.ServiceInterface) HandlerInterface {
	return &Handler{services: services}
}

func (h *Handler) GetUserByUsername(w http.ResponseWriter, r *http.Request) {
	username := httphelper.GetRouteParam(r, "username")
	user, err := h.services.GetUserByUsername(username)
	if err != nil {
		httphelper.JSONResponse(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	httphelper.JSONResponse(w, http.StatusOK, "User found", user)
}
