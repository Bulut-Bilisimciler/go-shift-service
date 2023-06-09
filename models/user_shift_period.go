package models

import (
	"time"
)

// UserShiftPeriod struct represents the shift period of a user
type UserShiftPeriod struct {
	ID        int64      `gorm:"primaryKey" json:"id"`
	UserID    int64      `gorm:"column:user_id" json:"userId"`
	StartDate time.Time  `gorm:"column:start_date" json:"startDate"`
	EndDate   time.Time  `gorm:"column:end_date" json:"endDate"`
	ShiftID   int64      `gorm:"column:shift_id" json:"shiftId"`
	CreatedAt time.Time  `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt time.Time  `gorm:"column:updated_at" json:"updatedAt"`
	DeletedAt *time.Time `gorm:"column:deleted_at" json:"-"`
}

// TableName returns the table name for the UserShiftPeriod model
func (UserShiftPeriod) TableName() string {
	return "user_shift_periods"
}
