package handlers

import (
	"fmt"
	"net/http"
	"user/services"

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

func (h *Handler) GetUserByUsername(w http.ResponseWriter, r *http.Request) {
	tracer := otel.Tracer(fmt.Sprintf("%s/handler", constant.ServiceUser))
	ctx, span := tracer.Start(r.Context(), "GetUserByUsername Handler")
	defer span.End()

	username := httphelper.GetQueryParams(r, "username")
	span.AddEvent("Checking user")
	user, err := h.services.GetUserByUsername(ctx, username)
	if err != nil {
		httphelper.JSONResponse(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	span.AddEvent("User found")
	httphelper.JSONResponse(w, http.StatusOK, "User found", user)
}
