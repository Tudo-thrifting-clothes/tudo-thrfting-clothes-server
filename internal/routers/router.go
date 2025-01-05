package routers

import (
	"tudo-thrfting-clothes-server/internal/controller"
	"tudo-thrfting-clothes-server/internal/middlewares"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	r.Use(middlewares.AuthMiddleware())

	path := r.Group("/v1")
	{
		path.GET("/ping/:id", controller.NewUserController().GetListuser)
	}

	return r
}
