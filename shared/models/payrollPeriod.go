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

type PayrollPeriodRequest struct {
	PeriodName string `json:"period_name"`
	StartDate  string `json:"start_date"`
	EndDate    string `json:"end_date"`
}

type PayrollPeriod struct {
	ID          uuid.UUID      `gorm:"type:uuid;primaryKey" json:"id"`
	PeriodName  string         `gorm:"size:50;not null" json:"period_name"`
	StartDate   time.Time      `gorm:"type:date;not null" json:"start_date"`
	EndDate     time.Time      `gorm:"type:date;not null" json:"end_date"`
	CreatedAt   time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	CreatedBy   uuid.UUID      `gorm:"type:uuid;not null" json:"created_by"`
	UpdatedBy   uuid.UUID      `gorm:"type:uuid" json:"updated_by"`
	OldDataJSON datatypes.JSON `gorm:"-" json:"-"`
}

func (p *PayrollPeriod) GetID() uuid.UUID {
	return p.ID
}

func (p *PayrollPeriod) GetTableName() string {
	return constant.PayrollPeriod
}

func (p *PayrollPeriod) GetOldData() datatypes.JSON {
	return p.OldDataJSON
}

func (p *PayrollPeriod) GetNewData() datatypes.JSON {
	data, _ := json.Marshal(p)
	return data
}

func (p *PayrollPeriod) BeforeCreate(tx *gorm.DB) (err error) {
	userID := tx.Statement.Context.Value(constant.ContextUserID).(string)
	uId, _ := utils.ParseUUID(userID)

	p.CreatedBy = uId
	p.ID = utils.GenerateUUID()
	return nil
}

func (p *PayrollPeriod) AfterCreate(tx *gorm.DB) (err error) {
	audit.CreateLog(tx, constant.Create, p)
	return nil
}
