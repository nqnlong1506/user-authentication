package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/nqnlong1506/user-authentication/controller"
	"github.com/nqnlong1506/user-authentication/database"
	"github.com/nqnlong1506/user-authentication/models"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	{
		database.ConnecetDB()
		models.InitializeUser()
	}
}

func main() {
	r := gin.Default()

	r.POST("/signup", controller.Signup)
	r.POST("/login", controller.Login)

	r.POST("/validate", controller.Validation)

	r.Use(controller.Validation)
	r.GET("/do-something", controller.DoSomeThing)

	r.Run() // listen and serve on 0.0.0.0:8080
}
