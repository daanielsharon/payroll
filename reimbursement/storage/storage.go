package storage

import (
	"context"
	shared_context "shared/context"
	"shared/models"
	"shared/utils"
	"time"

	"gorm.io/gorm"
)

type DB struct {
	DB *gorm.DB
}

func NewStorage(db *gorm.DB) Storage {
	return &DB{DB: db}
}

func (s *DB) Submit(ctx context.Context, reimbursement models.Reimbursement) error {
	userId, _ := shared_context.GetUserID(ctx)
	uId, _ := utils.ParseUUID(userId)
	reimbursement.UserID = uId

	return s.DB.WithContext(ctx).Create(&reimbursement).Error
}

func (s *DB) GetReimbursementByDate(ctx context.Context, startDate, endDate time.Time) []models.Reimbursement {
	var reimbursement []models.Reimbursement

	data := s.DB.WithContext(ctx).
		Where("date BETWEEN ? AND ?", startDate, endDate).
		Find(&reimbursement)

	if data.Error != nil {
		return nil
	}

	return reimbursement
}

func (s *DB) UpdatePayroll(ctx context.Context, reimbursement models.Reimbursement) error {
	return s.DB.WithContext(ctx).Save(&reimbursement).Error
}
