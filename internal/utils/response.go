package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func JSONResponse(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{"data": data})
}

func ErrorResponse(c *gin.Context, status int, message string, err error) {
	c.JSON(status, gin.H{
		"error":   message,
		"details": err.Error(),
	})
}
