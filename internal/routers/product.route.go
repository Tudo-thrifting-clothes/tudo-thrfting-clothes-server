package routers

import (
	"tudo-thrfting-clothes-server/internal/controller"

	"github.com/gin-gonic/gin"
)

type ProductRoute struct{}

func (ur *ProductRoute) InitUserRouter(Router *gin.RouterGroup) {

	uc := controller.NewProductController()

	userRouterPublic := Router.Group("/user")
	{
		userRouterPublic.GET("/", uc.GetListuser)
	}

	userRouterPrivate := Router.Group("/user")
	// userRouterPrivate.Use(LimitMiddleware())
	// userRouterPrivate.Use(AuthMiddleware())
	// userRouterPrivate.Use(PermissionMiddleware())
	{
		userRouterPrivate.GET("/get_info")
	}
}
