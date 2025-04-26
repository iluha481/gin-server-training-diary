package models

import "gorm.io/gorm"

type WorkoutPost struct {
	gorm.Model
	UserID      uint
	User        User
	Title       string `gorm:"not null"`
	Description string `gorm:"not null"`
	ImageURL    string
}
