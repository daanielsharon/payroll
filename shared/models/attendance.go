package models

import (
	"time"

	"github.com/google/uuid"
)

type Attendance struct {
	ID           uuid.UUID  `gorm:"type:uuid;primaryKey" json:"id"`
	UserID       uuid.UUID  `gorm:"type:uuid;not null;uniqueIndex:idx_user_date" json:"user_id"`
	Date         time.Time  `gorm:"type:date;not null;uniqueIndex:idx_user_date" json:"date"`
	ClockInAt    *time.Time `gorm:"type:timestamptz" json:"clock_in_at"`
	ClockOutAt   *time.Time `gorm:"type:timestamptz" json:"clock_out_at"`
	HoursWorked  *float64   `gorm:"type:numeric(4,2)" json:"hours_worked"` // e.g 7.50
	PayrollRunID *uuid.UUID `gorm:"type:uuid;index" json:"payroll_run_id"`
	CreatedAt    time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt    time.Time  `gorm:"autoUpdateTime" json:"updated_at"`
	CreatedBy    uuid.UUID  `gorm:"type:uuid;not null" json:"created_by"`
	UpdatedBy    uuid.UUID  `gorm:"type:uuid;not null" json:"updated_by"`

	User       User       `gorm:"foreignKey:UserID" json:"user"`
	PayrollRun PayrollRun `gorm:"foreignKey:PayrollRunID" json:"payroll_run"`
}
