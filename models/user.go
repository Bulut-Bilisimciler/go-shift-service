package models

import "time"

type User struct {
	ID         int64     `json:"id" gorm:"primary_key;not null"`
	UserID     int64     `json:"user_id" gorm:"column:user_id"` // user id from your organization
	Name       string    `json:"name" gorm:"column:name:"`
	Surname    string    `json:"surname" gorm:"column:surname:"`
	Email      string    `json:"email" gorm:"column:email:"`
	Made_Field string    `json:"made_field"  gorm:"column:made_field:"`
	Nickname   string    `json:"nickname"  gorm:"column:nickname:"`
	CreatedAt  time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"column:updated_at"`
}

// TableName overrides the table name used by User to `users`
func (u User) TableName() string {
	return "users"
}
