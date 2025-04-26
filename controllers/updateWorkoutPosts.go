package controllers

import (
	"net/http"
	"projectgin/initializers"
	"projectgin/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UpdateWorkoutPost(c *gin.Context) {
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
			"error": "workout not found",
		})
		return
	}

	if workoutPost.UserID != user.ID {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "user dont have permissions to update",
		})
		return
	}

	var body struct {
		Title       string
		Description string
		ImageURL    string
	}

	if c.BindJSON(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "bad request",
		})
		return
	}

	workoutPost.Title = body.Title
	workoutPost.Description = body.Description
	workoutPost.ImageURL = body.ImageURL

	initializers.DB.Save(&workoutPost)

	c.JSON(http.StatusOK, workoutPost)
}
