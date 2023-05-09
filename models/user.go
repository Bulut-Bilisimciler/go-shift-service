package models

import "time"

type User struct {
	Id        int64
	UserId    int64  `json:"user_id" gorm:"primary_key"` //sicilno
	Name      string `json:"name"`
	Surname   string `json:"surname"`
	Email     string `json:"mail"`
	isActive  bool
	CreatedAt *time.Time `json:"created_at" gorm:"default:now()"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"default:now()"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" gorm:"default:null"`
}
