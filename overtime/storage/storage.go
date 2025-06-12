package storage

import (
	"context"

	shared_context "shared/context"
	"shared/models"
	"shared/utils"

	"gorm.io/gorm"
)

type DB struct {
	DB *gorm.DB
}

func NewStorage(db *gorm.DB) Storage {
	return &DB{DB: db}
}

func (s *DB) Submit(ctx context.Context, overtime models.Overtime) error {
	userId, _ := shared_context.GetUserID(ctx)
	uId, _ := utils.ParseUUID(userId)
	overtime.UserID = uId
	return s.DB.WithContext(ctx).Create(&overtime).Error
}
