package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/nqnlong1506/user-authentication/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, %v", err)
	}

	router := echo.New()
	router.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	router.POST("/auth/login", routes.Login)
	router.POST("/auth/register", routes.Register)

	// test
	{
		router.POST("/nqnlong1506/test", test.Test)
	}

	router.Logger.Fatal(router.Start(":1323"))
}
