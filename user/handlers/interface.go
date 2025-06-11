package handlers

import "net/http"

type HandlerInterface interface {
	GetUserByUsername(w http.ResponseWriter, r *http.Request)
}
