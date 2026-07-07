package entity

import "time"

type User struct{
	ID 			uint		`gorm:"primary_key;column:id"`
	Fullname 	string		`gorm:"column:full_name"`
	Email 		string		`gorm:"column:email"`
	Password 	string		`gorm:"column:password"`
	role		string		`gorm:"column:role;type:enum('admin','member');default:'member'"`
	CreatedAt	time.Time	`gorm:"column:created_at;autoCreateTime"`
	UpdatedAt	time.Time	`gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
}