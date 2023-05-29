package models

import "time"

type Shift struct {
	ID        int64     `json:"id" gorm:"primary_key"`
	ShiftID   string    `json:"shift_id" gorm:"not null"`
	UserID    int64     `json:"user_id" gorm:"default:null"`
	StartTime time.Time `json:"start_time" gorm:"not null"`
	EndTime   time.Time `json:"end_time" gorm:"not null"`
	CreatedAt time.Time `json:"created_at" gorm:"default:now()"`
	UpdatedAt time.Time `json:"updated_at" gorm:"default:now()"`
	MadeField string    `json:"made_field" gorm:"default:null"`
}
