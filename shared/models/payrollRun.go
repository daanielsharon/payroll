package models

import (
	"time"

	"github.com/google/uuid"
)

type PayrollRun struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	PeriodID  uuid.UUID `gorm:"type:uuid;not null;index"`
	RanAt     time.Time `gorm:"autoCreateTime"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	CreatedBy uuid.UUID `gorm:"type:uuid;not null"`
	UpdatedBy uuid.UUID `gorm:"type:uuid;not null"`

	Period PayrollPeriod `gorm:"foreignKey:PeriodID"`
}
