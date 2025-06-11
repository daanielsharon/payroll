package handlers

import (
	"auth/services"
	"encoding/json"
	"net/http"
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

	// user, err := h.storage.GetUser(credentials.Username)
	// if err != nil {
	// 	http.Error(w, "invalid credentials", http.StatusUnauthorized)
	// 	return
	// }

	// if user.Password != credentials.Password {
	// 	http.Error(w, "invalid credentials", http.StatusUnauthorized)
	// 	return
	// }
}
