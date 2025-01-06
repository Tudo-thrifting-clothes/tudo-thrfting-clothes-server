package main

import (
	_ "tudo-thrfting-clothes-server/cmd/swag/docs"
	"tudo-thrfting-clothes-server/internal/initialize"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           API Documentation Tudo Thrifting Clothes Server
// @version         1.0.0
// @description     This is a sample server celler server.
// @termsOfService  tudo-thrfting-clothes-server

// @contact.name   TEAM TIPSGO
// @contact.url    tudo-thrfting-clothes-server
// @contact.email  baohc110902@gmail.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath /v1
// @schema http

func main() {
	r := initialize.Run()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":8080")
}
