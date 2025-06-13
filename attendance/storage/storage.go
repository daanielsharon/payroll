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

func (s *DB) GetAttendanceByUserIdAndDate(ctx context.Context, date time.Time) *models.Attendance {
	var attendance models.Attendance

	userID, _ := shared_context.GetUserID(ctx)

	data := s.DB.WithContext(ctx).First(&attendance, "user_id = ? AND date = ?", userID, date)

	if data.Error != nil {
		return nil
	}

	return &attendance
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

func (s *DB) ClockIn(ctx context.Context) error {
	userId, _ := shared_context.GetUserID(ctx)
	fmt.Println("user id", userId)
	uId, _ := utils.ParseUUID(userId)
	now := time.Now()

	data := s.DB.WithContext(ctx).Create(&models.Attendance{
		UserID:    uId,
		Date:      now,
		ClockInAt: &now,
	})

	if data.Error != nil {
		return data.Error
	}

	return nil
}

func (s *DB) ClockOut(ctx context.Context, previousAttendance *models.Attendance) error {
	return s.DB.WithContext(ctx).
		Save(previousAttendance).Error
}

func (s *DB) GetAttendanceByDate(ctx context.Context, startDate, endDate time.Time) []models.Attendance {
	var attendance []models.Attendance

	data := s.DB.WithContext(ctx).
		Where("date BETWEEN ? AND ? AND payroll_run_id is null", startDate, endDate).
		Find(&attendance)

	if data.Error != nil {
		return nil
	}

	return attendance
}

func (s *DB) UpdatePayroll(ctx context.Context, attendance models.Attendance) error {
	return s.DB.WithContext(ctx).
		Save(&attendance).Error
}
