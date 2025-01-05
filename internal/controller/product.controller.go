package controller

import (
	"tudo-thrfting-clothes-server/internal/service"
	"tudo-thrfting-clothes-server/pkg/response"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	userService *service.UserService
}

func NewProductController() *ProductController {
	return &ProductController{
		userService: service.NewUserService(),
	}
}

// Controller => Service => Repo => Model => DB
func (uc *ProductController) GetListuser(c *gin.Context) {
	response.SuccessReponse(c, 20001, uc.userService.GetListUser())
}
