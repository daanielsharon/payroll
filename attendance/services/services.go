package services

import (
	"attendance/storage"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"shared/constant"
	"shared/models"
	"shared/utils"
	"time"

	"github.com/google/uuid"
	"go.opentelemetry.io/otel"
)

type Service struct {
	storage storage.Storage
	NowFunc func() time.Time
}

func NewService(storage storage.Storage, nowFunc func() time.Time) ServiceInterface {
	return &Service{storage: storage, NowFunc: nowFunc}
}

func (s *Service) GetAttendanceByUserIdAndDate(ctx context.Context, date time.Time) *models.Attendance {
	tracer := otel.Tracer(fmt.Sprintf("%s/service", constant.ServiceAttendance))
	ctx, span := tracer.Start(ctx, "GetAttendanceByUserIdAndDate Service")
	defer span.End()

	return s.storage.GetAttendanceByUserIdAndDate(ctx, date)
}

func (s *Service) GetAttendanceByUserId(ctx context.Context) *models.Attendance {
	tracer := otel.Tracer(fmt.Sprintf("%s/service", constant.ServiceAttendance))
	ctx, span := tracer.Start(ctx, "GetAttendanceByUserId Service")
	defer span.End()

	return s.storage.GetAttendanceByUserId(ctx)
}

func (s *Service) Attend(ctx context.Context) error {
	tracer := otel.Tracer(fmt.Sprintf("%s/service", constant.ServiceAttendance))
	ctx, span := tracer.Start(ctx, "Attend Service")
	defer span.End()

	if utils.IsWeekend(s.NowFunc()) {
		span.AddEvent("Weekend validation")
		return errors.New("unable to submit on weekend")
	}

	span.AddEvent("Checking attendance")
	attendance := s.GetAttendanceByUserId(ctx)

	hasAttendance := attendance != nil

	if hasAttendance {
		hasClockIn := attendance.ClockInAt != nil
		hasClockOut := attendance.ClockOutAt != nil

		if hasClockIn && hasClockOut {
			return errors.New("user has clocked in and out")
		}

		if !hasClockIn && hasClockOut {
			return errors.New("user has clocked out without clocking in")
		}

		if hasClockIn {
			span.AddEvent("Clocking out")
			return s.ClockOut(ctx, attendance)
		}
	}

	span.AddEvent("Clocking in")
	return s.ClockIn(ctx)
}

func (s *Service) ClockIn(ctx context.Context) error {
	tracer := otel.Tracer(fmt.Sprintf("%s/service", constant.ServiceAttendance))
	ctx, span := tracer.Start(ctx, "clockIn Service")
	defer span.End()

	err := s.storage.ClockIn(ctx)
	if err != nil {
		return err
	}

	span.AddEvent("Clocking in")
	return nil
}

func (s *Service) ClockOut(ctx context.Context, previousAttendance *models.Attendance) error {
	now := time.Now()
	tracer := otel.Tracer(fmt.Sprintf("%s/service", constant.ServiceAttendance))
	ctx, span := tracer.Start(ctx, "clockOut Service")
	defer span.End()

	hours_worked := time.Since(*previousAttendance.ClockInAt).Hours()
	span.AddEvent("Clocking out")

	data, _ := json.Marshal(previousAttendance)
	previousAttendance.OldDataJSON = data

	previousAttendance.ClockOutAt = &now
	previousAttendance.HoursWorked = &hours_worked
	return s.storage.ClockOut(ctx, previousAttendance)
}

func (s *Service) GetAttendanceByDate(ctx context.Context, startDate, endDate time.Time) []models.Attendance {
	tracer := otel.Tracer(fmt.Sprintf("%s/service", constant.ServiceAttendance))
	ctx, span := tracer.Start(ctx, "GetAttendanceByDate Service")
	defer span.End()

	return s.storage.GetAttendanceByDate(ctx, startDate, endDate)
}

func (s *Service) UpdateAttendancePayroll(ctx context.Context, startDate, endDate time.Time, payrollRunId uuid.UUID) error {
	tracer := otel.Tracer(fmt.Sprintf("%s/service", constant.ServiceAttendance))
	ctx, span := tracer.Start(ctx, "UpdateAttendancePayroll Service")
	defer span.End()

	attendance := s.GetAttendanceByDate(ctx, startDate, endDate)

	if len(attendance) == 0 {
		return errors.New("no attendance found")
	}

	for _, att := range attendance {
		att.PayrollRunID = &payrollRunId
		s.storage.UpdatePayroll(ctx, att)
	}

	span.AddEvent("Updating payroll")

	return nil
}
