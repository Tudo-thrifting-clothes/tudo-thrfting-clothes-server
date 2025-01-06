package routers

type RouterGroup struct {
	UserRouter
	ProductRoute
}

var RouterGroupApp = new(RouterGroup)
