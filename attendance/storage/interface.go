package storage

import (
	"context"
	"shared/models"
)

type Storage interface {
	GetAttendanceByUserId(ctx context.Context) *models.Attendance
	ClockIn(ctx context.Context) error
	ClockOut(ctx context.Context, previousAttendance *models.Attendance) error
}
