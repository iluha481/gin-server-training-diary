package models

import "gorm.io/gorm"

type Exercise struct {
	gorm.Model
	WorkoutID uint `gorm:"not null"`
	Workout   Workout
	Name      string `gorm:"not null"`
	Sets      int
	Reps      int
	Weight    float64
	Notes     string `gorm:"type:text"`
}
