package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nqnlong1506/user-authentication/database"
	"github.com/nqnlong1506/user-authentication/models"
	"golang.org/x/crypto/bcrypt"
)

type Body struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Signup(c *gin.Context) {
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

	// var check map[string]string
	json.Unmarshal(rawData, &u)

	// validate username and password
	if u.Username == "" || u.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "username or password needed",
		})
		return
	}

	// check existing username
	var existingUser models.User
	result := database.DB.First(&existingUser, "username = ?", u.Username)
	if result.Error == nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Failed to create user",
		})
		return
	}

	// hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "Failed to create user",
		})

		return
	}

	// store to databse
	create := database.DB.Create(&models.User{
		Username: u.Username,
		Password: string(hashedPassword),
	})

	if create.Error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   create.Error.Error(),
			"message": "Failed to create user",
		})
	}

	// response
	c.JSON(200, gin.H{
		"message": "signup successfully",
	})
}
