package controllers

import (
	"net/http"
	"projectgin/initializers"
	"projectgin/models"

	"github.com/gin-gonic/gin"
)

func CreateWorkoutPost(c *gin.Context) {
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

	workoutPost := models.WorkoutPost{
		UserID:      user.ID,
		User:        user,
		Title:       body.Title,
		Description: body.Description,
		ImageURL:    body.ImageURL,
	}

	result := initializers.DB.Create(&workoutPost)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to create workout post",
		})
		return
	}
	c.JSON(http.StatusCreated, workoutPost)
}
