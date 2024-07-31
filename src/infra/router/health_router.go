package router

import (
	"github.com/gin-gonic/gin"
	"github.com/takuma123-type/go-api-study/interface/controller"
)

// healthRouter registers the health check route.
func healthRouter(r *gin.Engine) {
	r.GET("/health", controller.HealthCheck)
}
