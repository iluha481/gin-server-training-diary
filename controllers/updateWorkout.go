package controllers

import (
	"net/http"
	"projectgin/initializers"
	"projectgin/models"

	"github.com/gin-gonic/gin"
)

func UpdateWorkout(c *gin.Context) {
	// Структура для одного упражнения, которая соответствует JSON с фронтенда
	type exercise_ struct {
		ID     uint    `json:"ID"`
		Name   string  `json:"Name" binding:"required"`
		Sets   int     `json:"Sets" binding:"required"`
		Reps   int     `json:"Reps" binding:"required"`
		Weight float64 `json:"Weight"`
		Notes  string  `json:"Notes"`
	}

	// Структура для тела запроса, соответствующая JSON с фронтенда
	var body struct {
		WorkoutID uint        `json:"WorkoutID" binding:"required,gt=0"`
		Notes     string      `json:"Notes"`
		Exercises []exercise_ `json:"Exercises"`
	}

	// Парсим тело запроса
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "bad request",
		})
		return
	}

	// Проверяем авторизацию пользователя
	userInterface, ok := c.Get("user")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "failed extract user data",
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

	// Находим тренировку по ID
	var workout models.Workout
	if err := initializers.DB.First(&workout, body.WorkoutID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "workout not found",
		})
		return
	}

	// Проверяем, что тренировка принадлежит текущему пользователю
	if workout.UserID != user.ID {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "you do not have permission to modify this workout",
		})
		return
	}

	// Обновляем заметки тренировки, если они указаны
	if body.Notes != "" {
		workout.Notes = body.Notes
		if err := initializers.DB.Save(&workout).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "failed to update workout notes",
			})
			return
		}
	}

	// Если упражнения не указаны, просто возвращаем успешный ответ
	if len(body.Exercises) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"message": "workout updated successfully",
			"workout": workout,
		})
		return
	}

	// Обрабатываем список упражнений
	for _, ex := range body.Exercises {
		if ex.ID != 0 {
			// Обновление существующего упражнения
			var exercise models.Exercise
			if err := initializers.DB.First(&exercise, ex.ID).Error; err != nil {
				c.JSON(http.StatusNotFound, gin.H{
					"error": "exercise not found",
				})
				return
			}

			// Проверяем, что упражнение принадлежит указанной тренировке
			if exercise.WorkoutID != workout.ID {
				c.JSON(http.StatusForbidden, gin.H{
					"error": "exercise does not belong to this workout",
				})
				return
			}

			// Обновляем данные упражнения
			exercise.Name = ex.Name
			exercise.Sets = ex.Sets
			exercise.Reps = ex.Reps
			exercise.Weight = ex.Weight
			exercise.Notes = ex.Notes

			if err := initializers.DB.Save(&exercise).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": "failed to update exercise",
				})
				return
			}
		} else {
			// Создание нового упражнения
			newExercise := models.Exercise{
				WorkoutID: workout.ID,
				Name:      ex.Name,
				Sets:      ex.Sets,
				Reps:      ex.Reps,
				Weight:    ex.Weight,
				Notes:     ex.Notes,
			}

			if err := initializers.DB.Create(&newExercise).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": "failed to create exercise",
				})
				return
			}
		}
	}

	// Загружаем обновленные данные тренировки с упражнениями
	if err := initializers.DB.Preload("Exercises").First(&workout, workout.ID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "error loading related data after update",
		})
		return
	}

	// Возвращаем успешный ответ
	c.JSON(http.StatusOK, gin.H{
		"message": "workout updated successfully",
		"workout": workout,
	})
}
