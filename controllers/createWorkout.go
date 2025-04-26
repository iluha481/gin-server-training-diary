package controllers

import (
	"net/http"
	"projectgin/initializers"
	"projectgin/models"
	"time"

	"github.com/gin-gonic/gin"
)

// Структура для упражнений в запросе
type ExerciseInput struct {
	Name   string  `json:"Name"`
	Sets   int     `json:"Sets"`
	Reps   int     `json:"Reps"`
	Weight float64 `json:"Weight"`
	Notes  string  `json:"Notes"`
}

func CreateWorkout(c *gin.Context) {
	// Проверка авторизации пользователя
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

	// Новая структура тела запроса с упражнениями
	var body struct {
		Date      string          `json:"Date"`
		Notes     string          `json:"Notes"`
		Exercises []ExerciseInput `json:"Exercises"`
	}
	if c.BindJSON(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "bad request",
		})
		return
	}

	// Парсинг даты
	date, err := time.Parse("2006-01-02", body.Date)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid date format, use YYYY-MM-DD",
		})
		return
	}

	// Создание тренировки
	workout := models.Workout{
		UserID: user.ID,
		User:   user,
		Date:   date,
		Notes:  body.Notes,
	}
	result := initializers.DB.Create(&workout)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "error creating workout",
		})
		return
	}

	// Добавление упражнений, если они есть
	if len(body.Exercises) > 0 {
		for _, ex := range body.Exercises {
			exercise := models.Exercise{
				WorkoutID: workout.ID,
				Name:      ex.Name,
				Sets:      ex.Sets,
				Reps:      ex.Reps,
				Weight:    ex.Weight,
				Notes:     ex.Notes,
			}
			result = initializers.DB.Create(&exercise)
			if result.Error != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": "error creating exercise",
				})
				return
			}
		}
	}

	c.JSON(http.StatusCreated, workout)
}
