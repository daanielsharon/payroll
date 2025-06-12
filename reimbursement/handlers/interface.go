package handlers

import "net/http"

type HandlerInterface interface {
	Run(w http.ResponseWriter, r *http.Request)
}
