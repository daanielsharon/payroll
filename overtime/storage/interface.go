package storage

import (
	"context"
	"shared/models"
	"time"
)

type Storage interface {
	Submit(ctx context.Context, overtime models.Overtime) error
	GetOvertimeByDate(ctx context.Context, startDate, endDate time.Time) []models.Overtime
	UpdatePayroll(ctx context.Context, overtime models.Overtime) error
}
