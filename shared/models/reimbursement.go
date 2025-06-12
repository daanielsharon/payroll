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

type ReimbursementRequest struct {
	Amount      float64 `json:"amount"`
	Description string  `json:"description"`
}

type Reimbursement struct {
	ID           uuid.UUID      `gorm:"type:uuid;primaryKey" json:"id"`
	UserID       uuid.UUID      `gorm:"type:uuid;not null;index" json:"user_id"`
	Amount       float64        `gorm:"type:numeric(14,2);not null;check:amount >= 0" json:"amount"`
	Description  string         `gorm:"type:text" json:"description"`
	PayrollRunID *uuid.UUID     `gorm:"type:uuid;index" json:"payroll_run_id"`
	CreatedAt    time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt    time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	CreatedBy    uuid.UUID      `gorm:"type:uuid;not null" json:"created_by"`
	UpdatedBy    *uuid.UUID     `gorm:"type:uuid" json:"updated_by"`
	OldDataJSON  datatypes.JSON `gorm:"-" json:"-"`

	User       User        `gorm:"foreignKey:UserID"`
	PayrollRun *PayrollRun `gorm:"foreignKey:PayrollRunID"`
}

func (o *Reimbursement) GetID() uuid.UUID {
	return o.ID
}

func (o *Reimbursement) GetTableName() string {
	return constant.Reimbursement
}

func (o *Reimbursement) GetOldData() datatypes.JSON {
	return o.OldDataJSON
}

func (o *Reimbursement) GetNewData() datatypes.JSON {
	data, _ := json.Marshal(o)
	return data
}

func (o *Reimbursement) BeforeCreate(tx *gorm.DB) (err error) {
	userID := tx.Statement.Context.Value(constant.ContextUserID).(string)
	uId, _ := utils.ParseUUID(userID)
	o.UserID = uId
	o.CreatedBy = uId
	o.ID = utils.GenerateUUID()
	return nil
}

func (o *Reimbursement) AfterCreate(tx *gorm.DB) (err error) {
	audit.CreateLog(tx, constant.Create, o)
	return nil
}

func (o *Reimbursement) BeforeUpdate(tx *gorm.DB) (err error) {
	userID := tx.Statement.Context.Value(constant.ContextUserID).(string)
	uId, _ := utils.ParseUUID(userID)

	o.UpdatedBy = &uId
	return nil
}

func (o *Reimbursement) AfterUpdate(tx *gorm.DB) (err error) {
	audit.CreateLog(tx, constant.Update, o)
	return nil
}
