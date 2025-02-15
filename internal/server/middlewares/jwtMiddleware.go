package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func (m *Middlewares) JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader(authHeader)
		if header == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"errors": "the 'Authorization' header is missing"})
			return
		}

		parts := strings.Split(header, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"errors": "invalid token format"})
			return
		}
		if len(parts[1]) == 0 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"errors": "token is empty"})
			return
		}

		claims, err := m.tknMng.ParseClaims(parts[1])
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"errors": err.Error()})
			return
		}

		c.Set("user_id", (*claims)["sub"])
		c.Set("username", (*claims)["username"])
		c.Next()
	}
}
