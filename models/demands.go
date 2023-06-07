package models

import (
	"time"
)

// Demand struct represents a demand for a specific shift
type Demand struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	UserID    uint       `gorm:"column:user_id" json:"userId"`
	ShiftID   uint       `gorm:"column:shift_id" json:"shiftId"`
	StartDate time.Time  `gorm:"column:start_date" json:"startDate"`
	EndDate   time.Time  `gorm:"column:end_date" json:"endDate"`
	CreatedAt time.Time  `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt time.Time  `gorm:"column:updated_at" json:"updatedAt"`
	DeletedAt *time.Time `gorm:"column:deleted_at" json:"-"`
}

// TableName returns the table name for the Demand model
func (Demand) TableName() string {
	return "demands"
}
