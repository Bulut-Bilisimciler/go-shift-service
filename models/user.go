package models

import "time"

type User struct {
	ID         int64     `json:"id" gorm:"primary_key"`
	UserID     int64     `json:"user_id"  gorm:"default:null"`
	Name       string    `json:"name"  gorm:"default:null"`
	Surname    string    `json:"surname"  gorm:"default:null"`
	Email      string    `json:"email"  gorm:"default:null"`
	Made_Field string    `json:"made_field"  gorm:"default:null"`
	Nickname   string    `json:"nickname"  gorm:"default:null"`
	CreatedAt  time.Time `json:"created_at" gorm:"default:now()"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"default:now()"`
}

// TableName overrides the table name used by User to `users`
func (u User) TableName() string {
	return "users"
}
