package controllers

import (
	"net/http"
	"projectgin/initializers"
	"projectgin/models"

	"github.com/gin-gonic/gin"
)

func GetWorkout(c *gin.Context) {
	userInterface, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "user unathorized",
		})
		return
	}
	user, ok := userInterface.(models.User)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "error extracting user data",
		})
		return
	}

	// monthString := c.Query("month")
	// if monthString == "" {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"error": "empty query string",
	// 	})
	// 	return
	// }

	// month, err := time.Parse("2006-01", monthString)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"error": "cannot convert date",
	// 	})
	// 	return

	// }
	// startOfMonth := time.Date(month.Year(), month.Month(), 1, 0, 0, 0, 0, month.Location())
	// endOfMonth := startOfMonth.AddDate(0, 1, -1).Add(23*time.Hour + 59*time.Minute + 59*time.Second)
	//.Where("date BETWEEN ? AND ?", startOfMonth, endOfMonth)
	var workouts []models.Workout
	result := initializers.DB.Where("user_id = ?", user.ID).Preload("User").Preload("Exercises").Find(&workouts)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error loading workouts"})
		return
	}
	c.JSON(http.StatusOK, workouts)
}
