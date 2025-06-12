package storage

import (
	"context"
	"shared/models"
	"time"
)

type Storage interface {
	GetAttendanceByUserId(ctx context.Context) *models.Attendance
	ClockIn(ctx context.Context, time time.Time) error
	ClockOut(ctx context.Context, hours_worked float64, time time.Time) error
}
