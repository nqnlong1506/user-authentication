package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DoSomeThing(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "do something",
	})
}
