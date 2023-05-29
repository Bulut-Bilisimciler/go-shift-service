package models

import "time"

type User struct {
	Id            int64      `json:"id" gorm:"primary_key;not null"`
	Name          string     `json:"name" gorm:"column:name:"`
	Surname       string     `json:"surname" gorm:"column:surname:"`
	Email         string     `json:"email" gorm:"column:email:"`
	Made_Field    string     `json:"made_field"  gorm:"column:made_field:"`
	Nickname      string     `json:"nickname"  gorm:"column:nickname:"`
	UserId        int64      `json:"user_id" gorm:"column:user_id"` //sicilno
	Creation_Date *time.Time `gorm:"column:created_date";default:now()" json:"creation_date"`
	Update_Date   *time.Time `gorm:"column:updated_date";default:now()" json:"update_date"`
}

func (User) TableName() string { return "user" }
