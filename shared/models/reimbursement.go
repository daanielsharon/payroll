package models

import (
	"time"

	"github.com/google/uuid"
)

type Reimbursement struct {
	ID           uuid.UUID   `gorm:"type:uuid;primaryKey"`
	UserID       uuid.UUID   `gorm:"type:uuid;not null;index"`
	User         User        `gorm:"foreignKey:UserID"`
	Amount       float64     `gorm:"type:numeric(14,2);not null;check:amount >= 0"`
	Description  string      `gorm:"type:text"`
	PayrollRunID *uuid.UUID  `gorm:"type:uuid;index"`
	PayrollRun   *PayrollRun `gorm:"foreignKey:PayrollRunID"`
	CreatedAt    time.Time   `gorm:"autoCreateTime"`
	UpdatedAt    time.Time   `gorm:"autoUpdateTime"`
	CreatedBy    uuid.UUID   `gorm:"type:uuid;not null"`
	UpdatedBy    uuid.UUID   `gorm:"type:uuid;not null"`
}
