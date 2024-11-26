package routers

import (
	"github.com/gin-gonic/gin"

	"tudo-thrifting-server/controllers"
)

func SetupRouter() *gin.Engine {
	// Initialize Gin router
	router := gin.Default()

	// Define routes
	router.GET("/users", controllers.GetUsers)
	router.GET("/users/:id", controllers.GetUserByID)

	return router
}
