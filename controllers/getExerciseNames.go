package controllers

import (
	"net/http"
	"projectgin/initializers"
	"projectgin/models"

	"github.com/gin-gonic/gin"
)

func GetExerciseName(c *gin.Context) {

	var exercisesnames []string
	err := initializers.DB.Model(&models.ExerciseName{}).Distinct("name").Pluck("name", &exercisesnames).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "error getting list",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"exercises": exercisesnames,
	})
}
