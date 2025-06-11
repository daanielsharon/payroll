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

func (s *DB) GetUserByUsername(ctx context.Context, username string) (*models.User, error) {
	var user models.User
	if err := s.DB.WithContext(ctx).Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
