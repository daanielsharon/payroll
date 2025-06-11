package models

import (
	"time"

	"github.com/google/uuid"
)

type PayrollPeriod struct {
	ID         uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	PeriodName string    `gorm:"size:50;not null" json:"period_name"`
	StartDate  time.Time `gorm:"type:date;not null" json:"start_date"`
	EndDate    time.Time `gorm:"type:date;not null" json:"end_date"`
	CreatedAt  time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	CreatedBy  uuid.UUID `gorm:"type:uuid;not null" json:"created_by"`
	UpdatedBy  uuid.UUID `gorm:"type:uuid;not null" json:"updated_by"`
}
