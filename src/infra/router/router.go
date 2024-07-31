package router

import (
	"github.com/gin-gonic/gin"
)

// InitRouter initializes the router and sets up the routes.
func InitRouter() *gin.Engine {
	router := gin.Default()
	healthRouter(router) // Register the health check route
	// Other routers can be added here as needed
	return router
}
