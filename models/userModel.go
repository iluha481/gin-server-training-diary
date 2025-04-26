package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username  string `gorm:"unique"`
	Email     string `gorm:"unique" json:"-"`
	Password  string `json:"-"`
	AvatarURL string
}
