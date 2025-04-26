package controllers

import (
	"errors"
	"net/http"
	"projectgin/initializers"
	"projectgin/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DeleteExercise(c *gin.Context) {
	var body struct {
		ID uint `json:"ID"`
	}
	if c.BindJSON(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "bad request",
		})
		return
	}

	userInterface, ok := c.Get("user")
	if !ok {
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

	var exerciseStruct models.Exercise
	err := initializers.DB.Preload("Workout.User").First(&exerciseStruct, body.ID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "exercise not found",
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":   "database error",
				"details": err.Error(),
			})
		}
		return
	}

	if user.ID != exerciseStruct.Workout.User.ID {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "user dont have permission to delete",
		})
		return
	}

	initializers.DB.Delete(&exerciseStruct)
	c.JSON(http.StatusOK, gin.H{
		"message": "exercise deleted successfully",
	})
}
