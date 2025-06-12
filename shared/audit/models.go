package audit

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type Log struct {
	ID          uuid.UUID      `gorm:"type:uuid;primaryKey" json:"id"`
	TableName   string         `gorm:"size:50;not null;index" json:"table_name"`
	RecordID    uuid.UUID      `gorm:"type:uuid;not null;index" json:"record_id"`
	Action      string         `gorm:"size:20;not null;check:action IN ('create', 'update', 'delete')" json:"action"`
	PerformedBy *uuid.UUID     `gorm:"type:uuid;not null" json:"performed_by"`
	IPAddress   *string        `gorm:"type:text" json:"ip_address"`
	TraceID     *string        `gorm:"type:text;index" json:"trace_id"`
	RequestID   *uuid.UUID     `gorm:"type:uuid;index" json:"request_id"`
	OldData     datatypes.JSON `gorm:"type:jsonb" json:"old_data"`
	NewData     datatypes.JSON `gorm:"type:jsonb" json:"new_data"`
	CreatedAt   time.Time      `gorm:"autoCreateTime" json:"created_at"`
}
