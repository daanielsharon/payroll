package models

import (
	"time"

	"github.com/google/uuid"
)

type Attendance struct {
	ID           uuid.UUID  `gorm:"type:uuid;primaryKey"`
	UserID       uuid.UUID  `gorm:"type:uuid;not null;uniqueIndex:idx_attendance_user_date"`
	Date         time.Time  `gorm:"type:timestampz;not null;uniqueIndex:idx_attendance_user_date"`
	PayrollRunID *uuid.UUID `gorm:"type:uuid;index"`
	CreatedAt    time.Time  `gorm:"autoCreateTime"`
	UpdatedAt    time.Time  `gorm:"autoUpdateTime"`
	CreatedBy    uuid.UUID  `gorm:"type:uuid;not null"`
	UpdatedBy    uuid.UUID  `gorm:"type:uuid;not null"`

	PayrollRun PayrollRun `gorm:"foreignKey:PayrollRunID"`
	User       User       `gorm:"foreignKey:UserID"`
}
