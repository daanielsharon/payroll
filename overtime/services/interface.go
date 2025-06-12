package services

import (
	"context"
	"shared/models"
	"time"

	"github.com/google/uuid"
)

type ServiceInterface interface {
	Submit(ctx context.Context, overtime models.Overtime) error
	GetOvertimeByDate(ctx context.Context, startDate, endDate time.Time) []models.Overtime
	UpdateOvertimePayroll(ctx context.Context, startDate, endDate time.Time, payrollRunId uuid.UUID) error
}
