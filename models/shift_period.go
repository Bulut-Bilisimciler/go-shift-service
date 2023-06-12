package models

import (
	"time"
)

// UserShiftPeriod struct represents the shift period of a user
type ShiftPeriod struct {
	ID            int64     `json:"id" gorm:"primary_key;autoIncrement"`
	ShiftPeriodID int64     `json:"shift_period_id" gorm:"not null"`
	StartTime     time.Time `json:"start_time" gorm:"not null"`
	EndTime       time.Time `json:"end_time" gorm:"not null"`
	CreatedAt     time.Time `json:"created_at" gorm:"default:now()"`
	UpdatedAt     time.Time `json:"updated_at" gorm:"default:now()"`
}

// TableName returns the table name for the UserShiftPeriod model
func (ShiftPeriod) TableName() string {
	return "shift_periods"
}
