package controllers

import (
	"net/http"
	"projectgin/initializers"
	"projectgin/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DeleteWorkoutPost(c *gin.Context) {
	userInterface, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "user unauthorized",
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

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var workoutPost models.WorkoutPost
	result := initializers.DB.Preload("User").First(&workoutPost, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "post not found",
		})
		return
	}

	if workoutPost.UserID != user.ID {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "user dont have permissions to update",
		})
		return
	}

	initializers.DB.Delete(&workoutPost)
	c.JSON(http.StatusOK, gin.H{
		"message": "workout deleted",
	})
}
