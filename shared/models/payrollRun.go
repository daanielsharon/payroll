package models

import (
	"time"

	"github.com/google/uuid"
)

type PayrollRun struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	PeriodID  uuid.UUID `gorm:"type:uuid;not null;index" json:"period_id"`
	RanAt     time.Time `gorm:"autoCreateTime" json:"ran_at"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	CreatedBy uuid.UUID `gorm:"type:uuid;not null" json:"created_by"`
	UpdatedBy uuid.UUID `gorm:"type:uuid;" json:"updated_by"`

	Period PayrollPeriod `gorm:"foreignKey:PeriodID" json:"period"`
}
