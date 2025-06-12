package services

import (
	"context"
	"shared/models"
	"time"

	"github.com/google/uuid"
)

type ServiceInterface interface {
	GetAttendanceByUserId(ctx context.Context) *models.Attendance
	GetAttendanceByUserIdAndDate(ctx context.Context, date time.Time) *models.Attendance
	Attend(ctx context.Context) error
	UpdateAttendancePayroll(ctx context.Context, startDate, endDate time.Time, payrollRunId uuid.UUID) error
	ClockIn(ctx context.Context) error
	ClockOut(ctx context.Context, previousAttendance *models.Attendance) error
}
