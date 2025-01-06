package initialize

import (
	"tudo-thrfting-clothes-server/global"
	"tudo-thrfting-clothes-server/internal/routers"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	var r *gin.Engine
	if global.Config.Server.Enviroment == "dev" {
		gin.SetMode(gin.DebugMode)
		gin.ForceConsoleColor()
		r = gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
	}

	// r.Use()  //logging
	// r.Use() // cors
	// r.Use() // limit global request
	userRouter := routers.RouterGroupApp.UserRouter
	productRouter := routers.RouterGroupApp.ProductRoute

	mainGroup := r.Group("/v1")
	{
		mainGroup.GET("/ping")
	}
	{
		userRouter.InitUserRouter(mainGroup)
		productRouter.InitUserRouter(mainGroup)
	}

	return r
}
