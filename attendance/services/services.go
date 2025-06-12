package services

import (
	"attendance/storage"
	"context"
	"fmt"
	"shared/constant"
	"shared/models"
	"time"

	"go.opentelemetry.io/otel"
)

type Service struct {
	storage storage.Storage
}

func NewService(storage storage.Storage) ServiceInterface {
	return &Service{storage: storage}
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

	span.AddEvent("Checking attendance")
	attendance := s.GetAttendanceByUserId(ctx)

	if attendance != nil {
		span.AddEvent("Clocking in")
		return s.clockOut(ctx, attendance)
	}

	span.AddEvent("Clocking out")
	return s.clockIn(ctx)
}

func (s *Service) clockIn(ctx context.Context) error {
	tracer := otel.Tracer(fmt.Sprintf("%s/service", constant.ServiceAttendance))
	ctx, span := tracer.Start(ctx, "clockIn Service")
	defer span.End()

	err := s.storage.ClockIn(ctx, time.Now())
	if err != nil {
		return err
	}

	span.AddEvent("Clocking in")
	return nil
}

func (s *Service) clockOut(ctx context.Context, previousAttendance *models.Attendance) error {
	tracer := otel.Tracer(fmt.Sprintf("%s/service", constant.ServiceAttendance))
	ctx, span := tracer.Start(ctx, "clockOut Service")
	defer span.End()

	hours_worked := time.Since(*previousAttendance.ClockInAt).Hours()
	span.AddEvent("Clocking out")
	return s.storage.ClockOut(ctx, hours_worked, time.Now())
}
