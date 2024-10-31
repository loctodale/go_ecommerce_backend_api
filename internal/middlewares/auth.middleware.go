package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/loctodale/go-ecommerce-backend-api/pkg/reponse"
)

func AuthenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token != "valid-token" {
			reponse.ErrorResponse(c, reponse.ErrInvalidToken)
			c.Abort()
			return
		}
		c.Next()
	}
}
