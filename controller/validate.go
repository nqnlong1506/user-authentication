package controller

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/nqnlong1506/user-authentication/database"
	"github.com/nqnlong1506/user-authentication/models"
)

func Validation(c *gin.Context) {
	// get authorization
	tokenString, err := c.Cookie("Authorization")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	// parse token
	signingKey := []byte("lansfgsgasoiniwr")
	parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})

	if err != nil {
		fmt.Println("Failed to parse the token:", err)
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Failed to parse the token:" + err.Error(),
		})
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	claims := parsedToken.Claims.(jwt.MapClaims)

	// check existing user
	username := claims["username"]
	var existingUser models.User
	result := database.DB.First(&existingUser, "username = ?", username)
	if result.Error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "User does not exist",
		})
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	// reponse
	// c.JSON(http.StatusOK, gin.H{
	// 	"user": existingUser,
	// })

	// c.JSON(http.StatusOK, gin.H{})
	c.Next()
}
