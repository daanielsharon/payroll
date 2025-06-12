package handlers

import "net/http"

type HandlerInterface interface {
	Attendance(w http.ResponseWriter, r *http.Request)
}
