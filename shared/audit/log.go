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
	performedBy, _ := ctx.Value(constant.ContextUserID).(string)
	performedByUUID, _ := uuid.Parse(performedBy)
	ip, _ := ctx.Value(constant.ContextIPAddress).(string)
	requestID, _ := ctx.Value(constant.ContextRequestID).(string)
	requestIDUUID, _ := uuid.Parse(requestID)
	traceID, _ := ctx.Value(constant.ContextTraceID).(string)

	fmt.Println("OLD DATA", obj.GetOldData())
	log := Log{
		ID:          uuid.New(),
		TableName:   obj.GetTableName(),
		RecordID:    obj.GetID(),
		Action:      action,
		PerformedBy: &performedByUUID,
		IPAddress:   &ip,
		RequestID:   &requestIDUUID,
		TraceID:     &traceID,
		OldData:     obj.GetOldData(),
		NewData:     obj.GetNewData(),
	}

	return tx.Create(&log).Error
}
