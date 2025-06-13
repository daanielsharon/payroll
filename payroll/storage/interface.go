package storage

import (
	"context"
	"shared/models"

	"github.com/google/uuid"
)

type Storage interface {
	PayrollPeriodStorageInterface
	PayrollRunStorageInterface
}

type PayrollPeriodStorageInterface interface {
	CreatePayrollPeriod(ctx context.Context, payrollPeriod *models.PayrollPeriod) error
	GetPayrollPeriodById(ctx context.Context, id uuid.UUID) *models.PayrollPeriod
}

type PayrollRunStorageInterface interface {
	CreatePayrollRun(ctx context.Context, payrollRun *models.PayrollRun) error
}
