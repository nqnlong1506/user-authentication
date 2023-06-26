package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/nqnlong1506/user-authentication/database"
	"github.com/nqnlong1506/user-authentication/models"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) {
	// get raw data from request
	rawData, err := c.GetRawData()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	if string(rawData) == "" {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	var u Body
	err = json.Unmarshal(rawData, &u)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	// validate username and password
	if u.Username == "" || u.Password == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "username or password needed",
		})
		return
	}

	// check existing username
	var existingUser models.User
	result := database.DB.First(&existingUser, "username = ?", u.Username)
	if result.Error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "username or password wrong",
		})
		return
	}

	// compare hash password
	err = bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(u.Password))
	if err != nil {
		// Passwords match
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "username or password wrong",
		})
		return
	}

	// create jwt token
	claims := jwt.MapClaims{
		"username": existingUser.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(), // Expiration time
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	secretKey := os.Getenv("SECRET_KEY")
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		fmt.Println("Error generating JWT token:", err)
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Error generating JWT token:" + err.Error(),
		})
		return
	}

	// store jwt to cookies
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"message": "login successfully",
	})
}
