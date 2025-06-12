package services

import (
	"context"
	"shared/models"
	"time"
)

type ServiceInterface interface {
	GetAttendanceByUserId(ctx context.Context) *models.Attendance
	GetAttendanceByUserIdAndDate(ctx context.Context, date time.Time) *models.Attendance
	Attend(ctx context.Context) error
	ClockIn(ctx context.Context) error
	ClockOut(ctx context.Context, previousAttendance *models.Attendance) error
}
