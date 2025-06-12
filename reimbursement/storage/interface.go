package storage

import (
	"context"
	"shared/models"
	"time"
)

type Storage interface {
	Submit(ctx context.Context, reimbursement models.Reimbursement) error
	GetReimbursementByDate(ctx context.Context, startDate, endDate time.Time) []models.Reimbursement
	UpdatePayroll(ctx context.Context, reimbursement models.Reimbursement) error
}
