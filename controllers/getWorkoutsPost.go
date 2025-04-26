package controllers

import (
	"net/http"
	"projectgin/initializers"
	"projectgin/models"

	"github.com/gin-gonic/gin"
)

func GetWorkoutsPost(c *gin.Context) {
	var workoutPosts []models.WorkoutPost

	result := initializers.DB.Preload("User").Find(&workoutPosts)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to fetch workouts",
		})
		return
	}
	c.JSON(http.StatusOK, workoutPosts)
}
