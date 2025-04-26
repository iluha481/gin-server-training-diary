package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUserData(c *gin.Context) {
	// sample func thar respond with user data
	user, _ := c.Get("user")

	c.JSON(http.StatusOK, user)
}
