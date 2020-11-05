package middleware

import (
	"net/http"

	"github.com/dvillarruel93/SALES_DIEGO_VILLARRUEL/internal/ticket"
	"github.com/gin-gonic/gin"
)

const authToken = "1234asdf"

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("auth-token")

		if token != authToken {
			c.AbortWithStatusJSON(http.StatusUnauthorized, ticket.Error{Error:"invalid token"})
			return
		}

		c.Next()
	}
}
