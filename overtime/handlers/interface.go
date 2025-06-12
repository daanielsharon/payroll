package handlers

import "net/http"

type HandlerInterface interface {
	Submit(w http.ResponseWriter, r *http.Request)
	UpdateOvertimePayroll(w http.ResponseWriter, r *http.Request)
}
