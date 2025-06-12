package models

import (
	"encoding/json"
	"shared/audit"
	"shared/constant"
	"shared/utils"
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type OvertimeRequest struct {
	Date  string  `json:"date"`
	Hours float64 `json:"hours"`
}

type Overtime struct {
	ID           uuid.UUID  `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	UserID       uuid.UUID  `gorm:"type:uuid;not null;index" json:"user_id"`
	Date         time.Time  `gorm:"type:date;not null;index" json:"date"`
	Hours        float64    `gorm:"type:numeric(4,2);not null;check:hours > 0 AND hours <= 3" json:"hours"` // e.g 7.50
	PayrollRunID *uuid.UUID `gorm:"type:uuid;index" json:"payroll_run_id"`
	CreatedAt    time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt    time.Time  `gorm:"autoUpdateTime" json:"updated_at"`
	CreatedBy    uuid.UUID  `gorm:"type:uuid;not null" json:"created_by"`
	UpdatedBy    *uuid.UUID `gorm:"type:uuid" json:"updated_by"`

	User         User           `gorm:"foreignKey:UserID" json:"user"`
	PayrollRun   PayrollRun     `gorm:"foreignKey:PayrollRunID" json:"payroll_run"`
	OriginalData *Overtime      `gorm:"-" json:"-"`
	OldDataJSON  datatypes.JSON `gorm:"-" json:"-"`
}

func (o *Overtime) GetID() uuid.UUID {
	return o.ID
}

func (o *Overtime) GetTableName() string {
	return constant.Overtime
}

func (o *Overtime) GetOldData() datatypes.JSON {
	return o.OldDataJSON
}

func (o *Overtime) GetNewData() datatypes.JSON {
	data, _ := json.Marshal(o)
	return data
}

func (o *Overtime) BeforeCreate(tx *gorm.DB) (err error) {
	userID := tx.Statement.Context.Value(constant.ContextUserID).(string)
	uId, _ := utils.ParseUUID(userID)
	o.CreatedBy = uId
	o.ID = utils.GenerateUUID()
	return nil
}

func (o *Overtime) AfterCreate(tx *gorm.DB) (err error) {
	audit.CreateLog(tx, constant.Create, o)
	return nil
}
