package storage

import (
	"context"
	"time"

	"shared/models"

	"gorm.io/gorm"
)

type DB struct {
	DB *gorm.DB
}

func NewStorage(db *gorm.DB) Storage {
	return &DB{DB: db}
}

func (s *DB) Submit(ctx context.Context, overtime models.Overtime) error {
	return s.DB.WithContext(ctx).Create(&overtime).Error
}

func (s *DB) GetOvertimeByDate(ctx context.Context, startDate, endDate time.Time) []models.Overtime {
	var overtime []models.Overtime

	data := s.DB.WithContext(ctx).
		Where("date BETWEEN ? AND ? AND payroll_run_id is null", startDate, endDate).
		Find(&overtime)

	if data.Error != nil {
		return nil
	}

	return overtime
}

func (s *DB) UpdatePayroll(ctx context.Context, overtime models.Overtime) error {
	return s.DB.WithContext(ctx).Save(&overtime).Error
}
