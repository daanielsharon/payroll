package services

import (
	"context"
	"shared/models"
)

type ServiceInterface interface {
	PayloadPeriodServiceInterface
	PayloadRunServiceInterface
}

type PayloadPeriodServiceInterface interface {
	CreatePayrollPeriod(ctx context.Context, payrollPeriod models.PayrollPeriod) error
}

type PayloadRunServiceInterface interface {
	CreatePayrollRun(ctx context.Context, payrollRun models.PayrollRun) error
}
