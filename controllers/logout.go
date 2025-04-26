package controllers

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func Logout(c *gin.Context) {
	host := os.Getenv("HOST")
	c.SetCookie("Authorization", "", -1, "", host, false, false)
	c.JSON(http.StatusOK, gin.H{
		"message": "logout successful",
	})
}
