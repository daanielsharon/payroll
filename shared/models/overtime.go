package models

import (
	"time"

	"github.com/google/uuid"
)

type Overtime struct {
	ID           uuid.UUID  `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	UserID       uuid.UUID  `gorm:"type:uuid;not null;index" json:"user_id"`
	Date         time.Time  `gorm:"type:date;not null;index" json:"date"`
	Hours        float64    `gorm:"type:numeric(4,2);not null;check:hours > 0 AND hours <= 3" json:"hours"` // e.g 7.50
	PayrollRunID *uuid.UUID `gorm:"type:uuid;index" json:"payroll_run_id"`
	CreatedAt    time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt    time.Time  `gorm:"autoUpdateTime" json:"updated_at"`
	CreatedBy    uuid.UUID  `gorm:"type:uuid;not null" json:"created_by"`
	UpdatedBy    uuid.UUID  `gorm:"type:uuid;not null" json:"updated_by"`

	User       User       `gorm:"foreignKey:UserID" json:"user"`
	PayrollRun PayrollRun `gorm:"foreignKey:PayrollRunID" json:"payroll_run"`
}
