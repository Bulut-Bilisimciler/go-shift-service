package models

import "time"

type User struct {
	ID           int64     `json:"id" gorm:"primary_key"`
	UserID       int64     `json:"user_id"  gorm:"default:null"` // Employee ID UNIQUE
	Name         string    `json:"name"  gorm:"default:null"`
	Surname      string    `json:"surname"  gorm:"default:null"`
	Email        string    `json:"email"  gorm:"default:null"`
	Company      string    `json:"company"  gorm:"default:null"`
	Department   string    `json:"department"  gorm:"default:null"`
	Directorship string    `json:"directorship"  gorm:"default:null"`
	Made_Field   string    `json:"made_field"  gorm:"default:null"`
	Username     string    `json:"username"  gorm:"default:null"`
	Type         string    `json:"type"  gorm:"default:null"`
	CreatedAt    time.Time `json:"created_at" gorm:"default:now()"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"default:now()"`
}

// TableName overrides the table name used by User to `users`
func (u User) TableName() string {
	return "users"
}
