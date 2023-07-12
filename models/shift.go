package models

import "time"

type Shift struct {
	ID         int64     `json:"id" gorm:"primary_key;autoIncrement"`
	ShiftID    string    `json:"shift_id" gorm:"not null"`
	EmployeeID int64     `json:"employee_id" gorm:"unique_index"`
	StartTime  time.Time `json:"start_time" gorm:"not null"`
	EndTime    time.Time `json:"end_time" gorm:"not null"`
	CreatedAt  time.Time `json:"created_at" gorm:"default:now()"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"default:now()"`
	Group_Name string    `json:"group_name" gorm:"default:null"`
}
