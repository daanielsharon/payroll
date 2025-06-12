package handlers

import "net/http"

type HandlerInterface interface {
	GetAttendanceByUserIdAndDate(w http.ResponseWriter, r *http.Request)
	Submit(w http.ResponseWriter, r *http.Request)
	UpdateAttendancePayroll(w http.ResponseWriter, r *http.Request)
}
