package storage

import (
	"context"
	"shared/models"
	"time"
)

type Storage interface {
	GetAttendanceByUserIdAndDate(ctx context.Context, date time.Time) *models.Attendance
	GetAttendanceByUserId(ctx context.Context) *models.Attendance
	GetAttendanceByDate(ctx context.Context, startDate, endDate time.Time) []models.Attendance
	UpdatePayroll(ctx context.Context, attendance models.Attendance) error
	ClockIn(ctx context.Context) error
	ClockOut(ctx context.Context, previousAttendance *models.Attendance) error
}
