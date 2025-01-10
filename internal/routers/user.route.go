package routers

import (
	"tudo-thrfting-clothes-server/internal/controller"

	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (ur *UserRouter) InitUserRouter(Router *gin.RouterGroup) {

	uc := controller.NewUserController()

	userRouterPublic := Router.Group("/user")
	{
		userRouterPublic.GET("/", uc.GetListuser)
	}

	// userRouterPrivate := Router.Group("/user")
	// // userRouterPrivate.Use(LimitMiddleware())
	// // userRouterPrivate.Use(AuthMiddleware())
	// // userRouterPrivate.Use(PermissionMiddleware())
	// {
	// 	userRouterPrivate.GET("/get_info")
	// }
}
