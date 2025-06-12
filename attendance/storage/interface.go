package storage

import (
	"context"
	"shared/models"
	"time"
)

type Storage interface {
	GetAttendanceByUserIdAndDate(ctx context.Context, date time.Time) *models.Attendance
	GetAttendanceByUserId(ctx context.Context) *models.Attendance
	ClockIn(ctx context.Context) error
	ClockOut(ctx context.Context, previousAttendance *models.Attendance) error
}
