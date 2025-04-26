package controllers

import (
	"errors"
	"net/http"
	"projectgin/initializers"
	"projectgin/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DeleteWorkout(c *gin.Context) {
	// Структура для получения ID тренировки из запроса
	var body struct {
		ID uint `json:"ID"`
	}
	if c.BindJSON(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "bad request",
		})
		return
	}

	// Проверяем, авторизован ли пользователь
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

	// Находим тренировку с предварительной загрузкой пользователя
	var workoutStruct models.Workout
	err := initializers.DB.Preload("User").First(&workoutStruct, body.ID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "workout not found",
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":   "database error",
				"details": err.Error(),
			})
		}
		return
	}

	// Проверяем, принадлежит ли тренировка текущему пользователю
	if user.ID != workoutStruct.User.ID {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "user does not have permission to delete this workout",
		})
		return
	}

	// Удаляем все упражнения, связанные с тренировкой
	err = initializers.DB.Where("workout_id = ?", workoutStruct.ID).Delete(&models.Exercise{}).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "error deleting exercises",
			"details": err.Error(),
		})
		return
	}

	// Удаляем саму тренировку
	err = initializers.DB.Delete(&workoutStruct).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "error deleting workout",
			"details": err.Error(),
		})
		return
	}

	// Успешный ответ
	c.JSON(http.StatusOK, gin.H{
		"message": "workout and associated exercises deleted successfully",
	})
}
