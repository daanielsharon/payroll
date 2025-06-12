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

type PayrollRunRequest struct {
	PeriodID uuid.UUID `json:"period_id"`
}

type PayrollRun struct {
	ID        uuid.UUID  `gorm:"type:uuid;primaryKey" json:"id"`
	PeriodID  uuid.UUID  `gorm:"type:uuid;not null;index" json:"period_id"`
	RanAt     time.Time  `gorm:"autoCreateTime" json:"ran_at"`
	CreatedAt time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time  `gorm:"autoUpdateTime" json:"updated_at"`
	CreatedBy uuid.UUID  `gorm:"type:uuid;not null" json:"created_by"`
	UpdatedBy *uuid.UUID `gorm:"type:uuid" json:"updated_by"`

	Period      PayrollPeriod  `gorm:"foreignKey:PeriodID" json:"period"`
	OldDataJSON datatypes.JSON `gorm:"-" json:"-"`
}

func (p *PayrollRun) GetID() uuid.UUID {
	return p.ID
}

func (p *PayrollRun) GetTableName() string {
	return constant.PayrollRun
}

func (p *PayrollRun) GetOldData() datatypes.JSON {
	return p.OldDataJSON
}

func (p *PayrollRun) GetNewData() datatypes.JSON {
	data, _ := json.Marshal(p)
	return data
}

func (p *PayrollRun) BeforeCreate(tx *gorm.DB) (err error) {
	userID := tx.Statement.Context.Value(constant.ContextUserID).(string)
	uId, _ := utils.ParseUUID(userID)

	p.CreatedBy = uId
	p.ID = utils.GenerateUUID()
	return nil
}

func (p *PayrollRun) AfterCreate(tx *gorm.DB) (err error) {
	audit.CreateLog(tx, constant.Create, p)
	return nil
}
