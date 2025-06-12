package services

import (
	"context"
	"shared/models"
	"time"

	"github.com/google/uuid"
)

type ServiceInterface interface {
	Submit(ctx context.Context, reimbursement models.Reimbursement) error
	GetReimbursementByDate(ctx context.Context, startDate, endDate time.Time) []models.Reimbursement
	UpdateReimbursementPayroll(ctx context.Context, startDate, endDate time.Time, payrollRunId uuid.UUID) error
}
