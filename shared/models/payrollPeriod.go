package models

import (
	"time"

	"github.com/google/uuid"
)

type PayrollPeriod struct {
	ID         uuid.UUID `gorm:"type:uuid;primaryKey"`
	PeriodName string    `gorm:"size:50;not null"`
	StartDate  time.Time `gorm:"type:date;not null"`
	EndDate    time.Time `gorm:"type:date;not null"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime"`
	CreatedBy  uuid.UUID `gorm:"type:uuid;not null"`
	UpdatedBy  uuid.UUID `gorm:"type:uuid;not null"`
}
