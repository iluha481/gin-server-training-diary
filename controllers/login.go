package controllers

import (
	"net/http"
	"os"
	"projectgin/initializers"
	"projectgin/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) {
	// create struct that contains user's json
	var body struct {
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
	// struct for finded user
	var user models.User
	// search for user with the same email, if not found respond with 400
	initializers.DB.First(&user, "email = ?", body.Email)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email",
		})
		return
	}
	// if hashes are the same continue
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid password",
		})
		return
	}
	// create new jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})
	// sing it with secret key
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create token",
		})
		return
	}
	// create cookie with the token
	c.SetSameSite(http.SameSiteLaxMode)
	// todo:
	// убрать это, вынести в конфиг
	host := os.Getenv("HOST")
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", host, false, false)
	// respond with ok
	c.JSON(http.StatusOK, gin.H{})
}
