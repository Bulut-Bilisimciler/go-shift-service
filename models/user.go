package models

import "time"

type User struct {
	ID           int64     `json:"id" gorm:"primary_key;autoIncrement"`
	EmployeeID   int64     `json:"employee_id"  gorm:"unique_index"` // Employee ID UNIQUE
	Name         string    `json:"name"  gorm:"default:null"`
	Surname      string    `json:"surname"  gorm:"default:null"`
	Email        string    `json:"email"  gorm:"default:null"`
	Company      string    `json:"company"  gorm:"default:null"`
	Department   string    `json:"department"  gorm:"default:null"`
	Directorship string    `json:"directorship"  gorm:"default:null"`
	Group_Name   string    `json:"group_name"  gorm:"default:null"`
	Username     string    `json:"username"  gorm:"default:null"`
	Type         string    `json:"type"  gorm:"default:null"`
	CreatedAt    time.Time `json:"created_at" gorm:"default:now()"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"default:now()"`
}

// TableName overrides the table name used by User to `users`
func (u User) TableName() string {
	return "users"
}
