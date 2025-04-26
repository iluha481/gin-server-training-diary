package controllers

import (
	"fmt"
	"net/http"
	"os"
	"projectgin/initializers"
	"projectgin/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func RequireAuth(c *gin.Context) {
	// middleware for auth
	// getting cookies
	tokenString, err := c.Cookie("Authorization")
	//if err, abort
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	// parsing token
	// idk
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET")), nil
	})
	if err != nil || !token.Valid {

		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	// if all is ok continue
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// checking if cookie is expired
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		var user models.User
		initializers.DB.First(&user, claims["sub"])
		// if user not found abort
		if user.ID == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Set("user", user)

		c.Next()

	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}
