package storage

import (
	"context"

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
