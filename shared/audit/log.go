package audit

import (
	"fmt"
	"shared/constant"

	"github.com/google/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Auditable interface {
	GetID() uuid.UUID
	GetTableName() string
	GetOldData() datatypes.JSON
	GetNewData() datatypes.JSON
}

func CreateLog(tx *gorm.DB, action string, obj Auditable) error {
	ctx := tx.Statement.Context
	performedBy, _ := ctx.Value(constant.ContextUserID).(uuid.UUID)
	fmt.Println("performedBy", performedBy)
	ip, _ := ctx.Value(constant.ContextIPAddress).(string)
	fmt.Println("ip", ip)
	requestID, _ := ctx.Value(constant.ContextRequestID).(uuid.UUID)
	fmt.Println("requestID", requestID)
	traceID, _ := ctx.Value(constant.ContextTraceID).(string)
	fmt.Println("traceID", traceID)

	log := Log{
		ID:          uuid.New(),
		TableName:   obj.GetTableName(),
		RecordID:    obj.GetID(),
		Action:      action,
		PerformedBy: &performedBy,
		IPAddress:   &ip,
		RequestID:   &requestID,
		TraceID:     &traceID,
		OldData:     obj.GetOldData(),
		NewData:     obj.GetNewData(),
	}

	return tx.Create(&log).Error
}
