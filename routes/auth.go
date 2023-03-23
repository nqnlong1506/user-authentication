package routes

import (
	"net/http"

	"github.com/labstack/echo"
)

func Login(c echo.Context) error {
	return c.String(http.StatusOK, "Login!")
}

func Register(c echo.Context) error {
	return c.String(http.StatusOK, "Register!")
}
