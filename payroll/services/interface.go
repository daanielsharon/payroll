package services

import (
	"context"
	"shared/models"

	"github.com/google/uuid"
)

type ServiceInterface interface {
	PayloadPeriodServiceInterface
	PayloadRunServiceInterface
}

type PayloadPeriodServiceInterface interface {
	CreatePayrollPeriod(ctx context.Context, payrollPeriod *models.PayrollPeriod) error
	GetPayrollPeriodById(ctx context.Context, id uuid.UUID) *models.PayrollPeriod
}

type PayloadRunServiceInterface interface {
	CreatePayrollRun(ctx context.Context, payrollRun models.PayrollRun) error
}
