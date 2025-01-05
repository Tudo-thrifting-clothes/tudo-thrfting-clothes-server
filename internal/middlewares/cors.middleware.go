package middlewares

import (
	"tudo-thrfting-clothes-server/pkg/response"

	"github.com/gin-gonic/gin"
)

func CorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			response.ErrorReponse(c, response.ErrUnauthorized, "")
		}
		c.Next()
	}
}
