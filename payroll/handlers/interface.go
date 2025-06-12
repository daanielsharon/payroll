package handlers

import "net/http"

type HandlerInterface interface {
	PayloadPeriodHandlerInterface
	PayloadRunHandlerInterface
}

type PayloadPeriodHandlerInterface interface {
	CreatePayrollPeriod(w http.ResponseWriter, r *http.Request)
}

type PayloadRunHandlerInterface interface {
	CreatePayrollRun(w http.ResponseWriter, r *http.Request)
}
