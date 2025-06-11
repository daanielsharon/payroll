package models

import (
	"time"

	"github.com/google/uuid"
)

type Overtime struct {
	ID           uuid.UUID  `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	UserID       uuid.UUID  `gorm:"type:uuid;not null;index"`
	Date         time.Time  `gorm:"type:date;not null;index"`
	Hours        float64    `gorm:"type:numeric(4,2);not null;check:hours > 0 AND hours <= 3"` // e.g 7.50
	PayrollRunID *uuid.UUID `gorm:"type:uuid;index"`
	CreatedAt    time.Time  `gorm:"autoCreateTime"`
	UpdatedAt    time.Time  `gorm:"autoUpdateTime"`
	CreatedBy    uuid.UUID  `gorm:"type:uuid;not null"`
	UpdatedBy    uuid.UUID  `gorm:"type:uuid;not null"`

	User       User       `gorm:"foreignKey:UserID"`
	PayrollRun PayrollRun `gorm:"foreignKey:PayrollRunID"`
}
