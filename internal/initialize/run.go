package initialize

import "github.com/gin-gonic/gin"

func Run() *gin.Engine {
	// Initialize the configuration
	InitConfig()
	// Initialize the logger
	InitLogger()
	// Initialize the database
	InitMySQL()
	// Initialize the Redis client
	InitRedis()

	r := InitRouter()

	return r
}
