package controllers

import (
	"net/http"
	"projectgin/initializers"
	"projectgin/models"

	"github.com/gin-gonic/gin"
)

func UpdateUserData(c *gin.Context) {

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
		Username  string
		AvatarURL string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "bad request",
		})
		return
	}
	user.Username = body.Username
	user.AvatarURL = body.AvatarURL
	initializers.DB.Save(user)

	c.JSON(http.StatusOK, gin.H{"message": "successful"})
}
