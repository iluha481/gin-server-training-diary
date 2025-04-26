package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	"projectgin/initializers"
	"projectgin/models"
)

func Signup(c *gin.Context) {
	// create struct that contains user's json
	var body struct {
		Username string
		Email    string
		Password string
	}
	// if json is bad
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Failed to read body",
		})
		return
	}
	// generate hash for user's password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	// if err
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Failed to hash password",
		})
		return
	}
	// create user struct with user's data
	user := models.User{Username: body.Username, Email: body.Email, Password: string(hash)}
	// create record in database
	result := initializers.DB.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Failed to create user",
		})
		return
	}
	// respond with OK
	c.JSON(http.StatusCreated, gin.H{})

}
