package services

import (
	"context"
	"shared/models"
)

type ServiceInterface interface {
	GetAttendanceByUserId(ctx context.Context) *models.Attendance
	Attend(ctx context.Context) error
	ClockIn(ctx context.Context) error
	ClockOut(ctx context.Context, previousAttendance *models.Attendance) error
}
