package handlers

import "net/http"

type HandlerInterface interface {
	Login(w http.ResponseWriter, r *http.Request)
}
