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

type Attendance struct {
	ID           uuid.UUID  `gorm:"type:uuid;primaryKey" json:"id"`
	UserID       uuid.UUID  `gorm:"type:uuid;not null;uniqueIndex:idx_user_date" json:"user_id"`
	Date         time.Time  `gorm:"type:date;not null;uniqueIndex:idx_user_date" json:"date"`
	ClockInAt    *time.Time `gorm:"type:timestamptz" json:"clock_in_at"`
	ClockOutAt   *time.Time `gorm:"type:timestamptz" json:"clock_out_at"`
	HoursWorked  *float64   `gorm:"type:numeric(4,2)" json:"hours_worked"` // e.g 7.50
	PayrollRunID *uuid.UUID `gorm:"type:uuid;index" json:"payroll_run_id"`
	CreatedAt    time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt    time.Time  `gorm:"autoUpdateTime" json:"updated_at"`
	CreatedBy    uuid.UUID  `gorm:"type:uuid;not null" json:"created_by"`
	UpdatedBy    *uuid.UUID `gorm:"type:uuid" json:"updated_by"`

	User        User           `gorm:"foreignKey:UserID" json:"user"`
	PayrollRun  PayrollRun     `gorm:"foreignKey:PayrollRunID" json:"payroll_run"`
	OldDataJSON datatypes.JSON `gorm:"-" json:"-"`
}

func (a *Attendance) GetID() uuid.UUID {
	return a.ID
}

func (a *Attendance) GetTableName() string {
	return constant.Attendance
}

func (a *Attendance) GetOldData() datatypes.JSON {
	return a.OldDataJSON
}

func (a *Attendance) GetNewData() datatypes.JSON {
	data, _ := json.Marshal(a)
	return data
}

func (a *Attendance) BeforeCreate(tx *gorm.DB) (err error) {
	userID := tx.Statement.Context.Value(constant.ContextUserID).(string)
	uId, _ := utils.ParseUUID(userID)

	a.CreatedBy = uId
	a.ID = utils.GenerateUUID()
	return nil
}

func (a *Attendance) AfterCreate(tx *gorm.DB) (err error) {
	audit.CreateLog(tx, constant.Create, a)
	return nil
}

func (a *Attendance) BeforeUpdate(tx *gorm.DB) (err error) {
	userID := tx.Statement.Context.Value(constant.ContextUserID).(string)
	uId, _ := utils.ParseUUID(userID)

	a.UpdatedBy = &uId
	return nil
}

func (a *Attendance) AfterUpdate(tx *gorm.DB) (err error) {
	audit.CreateLog(tx, constant.Update, a)
	return nil
}
