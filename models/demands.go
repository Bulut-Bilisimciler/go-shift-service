package models

import (
	"time"
)

// Demand struct represents a demand for a specific shift
type Demand struct {
	ID         int64      `json:"id" gorm:"primary_key;autoIncrement"`
	DemandID   int64      `json:"demand_id" gorm:"not null"`
	ShiftID    string     `json:"shift_id" gorm:"not null"`
	EmployeeID int64      `json:"employee_id" gorm:"unique_index"`
	StartTime  time.Time  `json:"start_time" gorm:"not null"`
	EndTime    time.Time  `json:"end_time" gorm:"not null"`
	CreatedAt  time.Time  `json:"created_at" gorm:"default:now()"`
	UpdatedAt  time.Time  `json:"updated_at" gorm:"default:now()"`
	DeletedAt  *time.Time `json:"deleted_at" gorm:"default:null"`
}

// TableName returns the table name for the Demand model
func (Demand) TableName() string {
	return "demands"
}
