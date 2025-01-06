package routers

import (
	"tudo-thrfting-clothes-server/internal/controller"

	"github.com/gin-gonic/gin"
)

type ProductRoute struct{}

func (ur *ProductRoute) InitUserRouter(Router *gin.RouterGroup) {

	pc := controller.NewProductController()

	productPublicRoute := Router.Group("/product")
	{
		productPublicRoute.GET("/", pc.GetListProduct)
	}

	userRouterPrivate := Router.Group("/product")
	// userRouterPrivate.Use(LimitMiddleware())
	// userRouterPrivate.Use(AuthMiddleware())
	// userRouterPrivate.Use(PermissionMiddleware())
	{
		userRouterPrivate.POST("/", pc.CreateProduct)
	}
}
