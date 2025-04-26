package controllers

import (
	"net/http"
	"projectgin/initializers"
	"projectgin/models"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	username := c.Query("username")
	if username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Empty user string"})
		return
	}
	var user models.User
	result := initializers.DB.Where("username = ?", username).First(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}
