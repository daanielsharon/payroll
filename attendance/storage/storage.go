package storage

import (
	"context"
	"fmt"
	"time"

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

func (s *DB) GetAttendanceByUserId(ctx context.Context) *models.Attendance {
	var attendance models.Attendance

	userID, _ := shared_context.GetUserID(ctx)
	data := s.DB.WithContext(ctx).First(&attendance, "user_id = ?", userID)

	if data.Error != nil {
		return nil
	}

	return &attendance
}

func (s *DB) ClockIn(ctx context.Context, time time.Time) error {
	userId, _ := shared_context.GetUserID(ctx)
	fmt.Println("user id", userId)
	uId, _ := utils.ParseUUID(userId)
	data := s.DB.WithContext(ctx).Create(&models.Attendance{
		UserID:    uId,
		Date:      time,
		ClockInAt: &time,
	})

	if data.Error != nil {
		return data.Error
	}

	return nil
}

func (s *DB) ClockOut(ctx context.Context, hours_worked float64, time time.Time) error {
	userId, _ := shared_context.GetUserID(ctx)
	data := s.DB.WithContext(ctx).
		Model(&models.Attendance{}).
		Where("user_id = ? AND date = ?", userId, time).
		Updates(&models.Attendance{
			ClockOutAt:  &time,
			HoursWorked: &hours_worked,
		})

	if data.Error != nil {
		return data.Error
	}

	return nil
}
