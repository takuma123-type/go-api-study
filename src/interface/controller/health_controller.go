package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HealthCheck handles the health check request.
func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}
