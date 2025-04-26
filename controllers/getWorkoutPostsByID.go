package controllers

import (
	"net/http"
	"projectgin/initializers"
	"projectgin/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetWorkoutPostsByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var workoutPost models.WorkoutPost
	if err := initializers.DB.Preload("User").First(&workoutPost, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Workout not found"})
		return
	}
	c.JSON(http.StatusOK, workoutPost)
}
