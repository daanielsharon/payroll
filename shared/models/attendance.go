package models

import (
	"time"

	"github.com/google/uuid"
)

type Attendance struct {
	ID           uuid.UUID  `gorm:"type:uuid;primaryKey"`
	UserID       uuid.UUID  `gorm:"type:uuid;not null;uniqueIndex:idx_user_date"`
	Date         time.Time  `gorm:"type:date;not null;uniqueIndex:idx_user_date"`
	ClockInAt    *time.Time `gorm:"type:timestamptz"`
	ClockOutAt   *time.Time `gorm:"type:timestamptz"`
	HoursWorked  *float64   `gorm:"type:numeric(4,2)"` // e.g 7.50
	PayrollRunID *uuid.UUID `gorm:"type:uuid;index"`
	CreatedAt    time.Time  `gorm:"autoCreateTime"`
	UpdatedAt    time.Time  `gorm:"autoUpdateTime"`
	CreatedBy    uuid.UUID  `gorm:"type:uuid;not null"`
	UpdatedBy    uuid.UUID  `gorm:"type:uuid;not null"`

	User       User       `gorm:"foreignKey:UserID"`
	PayrollRun PayrollRun `gorm:"foreignKey:PayrollRunID"`
}
