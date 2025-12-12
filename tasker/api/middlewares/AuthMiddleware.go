package middlewares

import (
	"github.com/Nebuska/neblab/shared/jwtAuth"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(manager *jwtAuth.JWTManager) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code":  http.StatusUnauthorized,
				"error": "No Authorization token",
			})
			return
		}

		token := strings.Split(authHeader, " ")
		if len(token) != 2 && strings.ToLower(token[0]) != "bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code":  http.StatusUnauthorized,
				"error": "Faulty Authorization token",
			})
			return
		}

		claims, err := manager.Verify(jwtAuth.JWTToken(token[1]))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code":  http.StatusUnauthorized,
				"error": err.Error(),
			})
			return
		}

		c.Set("claims", claims)
		c.Next()
	}
}
