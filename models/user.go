package models

import "time"

type User struct {
<<<<<<< HEAD
	Id            int64      `json:"id" gorm:"primary_key;not null"`
	Name          string     `json:"name" gorm:"column:name:"`
	Surname       string     `json:"surname" gorm:"column:surname:"`
	Email         string     `json:"email" gorm:"column:email:"`
	Made_Field    string     `json:"made_field"  gorm:"column:made_field:"`
	Nickname      string     `json:"nickname"  gorm:"column:nickname:"`
	UserId        int64      `json:"user_id" gorm:"column:user_id"` //sicilno
	Creation_Date *time.Time `gorm:"column:created_date";default:now()" json:"creation_date"`
	Update_Date   *time.Time `gorm:"column:updated_date";default:now()" json:"update_date"`
=======
	Id            int64      `json:"user_id" gorm:"primary_key;not null"`
	UserId        int64      `json:"user_id" gorm:"column:user_id"` //sicilno
	Name          string     `json:"name" gorm:"column:name:"`
	Surname       string     `json:"surname" gorm:"column:surname:"`
	Email         string     `json:"email" gorm:"column:email:"`
	Nickname      string     `json:"nickname"  gorm:"column:nickname:"`
	Creation_Date *time.Time `gorm:"column:creation_date";default:now()" json:"creation_date"`
	Update_Date   *time.Time `gorm:"column:update_date";default:now()" json:"update_date"`
	Made_Field    string     `json:"made_field"  gorm:"column:made_field:"`
>>>>>>> babb95df0fea093f40084966abceab22a79853b9
}

func (User) TableName() string { return "user" }
