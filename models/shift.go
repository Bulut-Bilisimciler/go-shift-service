package models

import "time"

type Shift struct {
	ShiftId   int64     `json:"shift_id" gorm:"primary_key"`
	UserId    int64     `json:"user_id" gorm:"foreign_key"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
	isActive  bool
	CreatedAt *time.Time `json:"created_at" gorm:"default:now()"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"default:now()"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" gorm:"default:null"`
}
