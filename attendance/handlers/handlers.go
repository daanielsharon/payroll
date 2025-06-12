package handlers

import (
	"attendance/services"
	"fmt"
	"log"
	"net/http"
	"shared/constant"
	httphelper "shared/http"
	"shared/models"
	"shared/utils"

	"go.opentelemetry.io/otel"
)

type Handler struct {
	services services.ServiceInterface
}

func NewHandler(services services.ServiceInterface) HandlerInterface {
	return &Handler{services: services}
}

func (h *Handler) Submit(w http.ResponseWriter, r *http.Request) {
	tracer := otel.Tracer(fmt.Sprintf("%s/handler", constant.ServiceAttendance))
	ctx, span := tracer.Start(r.Context(), "Submit Handler")
	defer span.End()

	span.AddEvent("Submitting")
	err := h.services.Attend(ctx)
	if err != nil {
		httphelper.JSONResponse(w, http.StatusInternalServerError, "Attendance run failed", nil)

		log.Println("Attendance run failed", err)
		return
	}

	httphelper.JSONResponse(w, http.StatusOK, "Attendance run successful", nil)
}

func (h *Handler) GetAttendanceByUserIdAndDate(w http.ResponseWriter, r *http.Request) {
	tracer := otel.Tracer(fmt.Sprintf("%s/handler", constant.ServiceAttendance))
	ctx, span := tracer.Start(r.Context(), "GetAttendanceByUserIdAndDate Handler")
	defer span.End()

	span.AddEvent("Getting attendance by user id and date")
	date := httphelper.GetQueryParams(r, "date")
	dateTime := utils.ConvertStringToDate(date)

	attendance := h.services.GetAttendanceByUserIdAndDate(ctx, dateTime)
	if attendance == nil {
		httphelper.JSONResponse(w, http.StatusNotFound, "Attendance not found", models.Attendance{})
		return
	}

	httphelper.JSONResponse(w, http.StatusOK, "Attendance found", attendance)
}

func (h *Handler) UpdateAttendancePayroll(w http.ResponseWriter, r *http.Request) {
	tracer := otel.Tracer(fmt.Sprintf("%s/handler", constant.ServiceAttendance))
	ctx, span := tracer.Start(r.Context(), "UpdateAttendancePayroll Handler")
	defer span.End()

	startDate := httphelper.GetQueryParams(r, "startDate")
	endDate := httphelper.GetQueryParams(r, "endDate")
	payrollRunId := httphelper.GetQueryParams(r, "payrollRunId")

	startDateParsed := utils.ConvertStringToDate(startDate)
	endDateParsed := utils.ConvertStringToDate(endDate)
	payrollRunIdParsed, _ := utils.ParseUUID(payrollRunId)

	span.AddEvent("Updating payroll")
	err := h.services.UpdateAttendancePayroll(ctx, startDateParsed, endDateParsed, payrollRunIdParsed)
	if err != nil {
		httphelper.JSONResponse(w, http.StatusInternalServerError, "Attendance payroll update failed", nil)

		log.Println("Attendance payroll update failed", err)
		return
	}

	httphelper.JSONResponse(w, http.StatusOK, "Attendance payroll update successful", nil)
}
