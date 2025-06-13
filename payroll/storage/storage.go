package storage

import (
	"context"
	"shared/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DB struct {
	DB *gorm.DB
}

func NewStorage(db *gorm.DB) Storage {
	return &DB{DB: db}
}

func (s *DB) CreatePayrollPeriod(ctx context.Context, payrollPeriod *models.PayrollPeriod) error {
	return s.DB.WithContext(ctx).Create(payrollPeriod).Error
}

func (s *DB) GetPayrollPeriodById(ctx context.Context, id uuid.UUID) *models.PayrollPeriod {
	var payrollPeriod models.PayrollPeriod

	data := s.DB.WithContext(ctx).First(&payrollPeriod, "id = ?", id)

	if data.Error != nil {
		return nil
	}

	return &payrollPeriod
}

func (s *DB) CreatePayrollRun(ctx context.Context, payrollRun *models.PayrollRun) error {
	return s.DB.WithContext(ctx).Create(payrollRun).Error
}
