package initializers

import (
	"log"
	"os"
	"projectgin/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	// initialize database
	dsn := os.Getenv("DB")
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err, dsn)
	}
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Exercise{})
	DB.AutoMigrate(&models.Workout{})
	DB.AutoMigrate(&models.ExerciseName{})
	DB.AutoMigrate(&models.WorkoutPost{})
}
