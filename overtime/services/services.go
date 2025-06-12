package services

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"overtime/storage"
	"shared/constant"
	httphelper "shared/http"
	"shared/models"
	"shared/utils"
	"time"

	"github.com/google/uuid"
	"go.opentelemetry.io/otel"
)

type Service struct {
	storage storage.Storage
	client  *httphelper.Client
}

func NewService(storage storage.Storage) ServiceInterface {
	return &Service{storage: storage, client: httphelper.NewWithServices()}
}

func (s *Service) Submit(ctx context.Context, overtime models.Overtime) error {
	tracer := otel.Tracer(fmt.Sprintf("%s/service", constant.ServiceOvertime))
	ctx, span := tracer.Start(ctx, "Submit Service")
	defer span.End()

	if overtime.Hours > 3 {
		return errors.New("hours must be less than or equal to 3")
	}

	if !utils.IsWeekend(overtime.Date) {
		span.AddEvent("Checking attendance")

		fmt.Println("over time date", overtime.Date)

		result, err := httphelper.DoAndDecode[models.Attendance](s.client, ctx, constant.ServiceAttendance, http.MethodGet, "/attendance?date="+overtime.Date.Format("2006-01-02"), nil)
		if err != nil {
			return err
		}

		hasClockIn := result.Data.ClockInAt != nil
		hasClockOut := result.Data.ClockOutAt != nil
		hasAttendance := hasClockIn && hasClockOut

		if !hasAttendance {
			return errors.New("user has not clocked in and out")
		}
	}

	return s.storage.Submit(ctx, overtime)
}

func (s *Service) GetOvertimeByDate(ctx context.Context, startDate, endDate time.Time) []models.Overtime {
	tracer := otel.Tracer(fmt.Sprintf("%s/service", constant.ServiceOvertime))
	ctx, span := tracer.Start(ctx, "GetOvertimeByDate Service")
	defer span.End()

	return s.storage.GetOvertimeByDate(ctx, startDate, endDate)
}

func (s *Service) UpdateOvertimePayroll(ctx context.Context, startDate, endDate time.Time, payrollRunId uuid.UUID) error {
	tracer := otel.Tracer(fmt.Sprintf("%s/service", constant.ServiceOvertime))
	ctx, span := tracer.Start(ctx, "UpdateOvertimePayroll Service")
	defer span.End()

	overtime := s.GetOvertimeByDate(ctx, startDate, endDate)
	if len(overtime) == 0 {
		return errors.New("no overtime found")
	}

	span.AddEvent("Updating payroll")

	for _, ot := range overtime {
		ot.PayrollRunID = &payrollRunId
		s.storage.UpdatePayroll(ctx, ot)
	}

	return nil
}
