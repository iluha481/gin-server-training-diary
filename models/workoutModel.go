package models

import (
	"time"

	"gorm.io/gorm"
)

type Workout struct {
	gorm.Model
	UserID    uint `gorm:"not null"`
	User      User
	Date      time.Time  `gorm:"type:timestamp;not null"`
	Notes     string     `gorm:"type:text"`
	Exercises []Exercise `gorm:"foreignKey:WorkoutID"`
}
