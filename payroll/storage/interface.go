package storage

import (
	"context"
	"shared/models"
)

type Storage interface {
	PayrollPeriodStorageInterface
	PayrollRunStorageInterface
}

type PayrollPeriodStorageInterface interface {
	CreatePayrollPeriod(ctx context.Context, payrollPeriod models.PayrollPeriod) error
}

type PayrollRunStorageInterface interface {
	CreatePayrollRun(ctx context.Context, payrollRun models.PayrollRun) error
}
