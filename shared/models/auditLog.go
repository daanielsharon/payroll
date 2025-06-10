package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type AuditLog struct {
	ID          uuid.UUID      `gorm:"type:uuid;primaryKey"`
	TableName   string         `gorm:"size:50;not null;index"`
	RecordID    uuid.UUID      `gorm:"type:uuid;not null;index"`
	Action      string         `gorm:"size:20;not null;check:action IN ('create', 'update', 'delete')"`
	PerformedBy uuid.UUID      `gorm:"type:uuid;not null"`
	IPAddress   string         `gorm:"type:inet"`
	RequestID   uuid.UUID      `gorm:"type:uuid;index"`
	OldData     datatypes.JSON `gorm:"type:jsonb"`
	NewData     datatypes.JSON `gorm:"type:jsonb"`
	CreatedAt   time.Time      `gorm:"autoCreateTime"`
}
