package middlewares

import (
	"github.com/Nebuska/neblab/shared/logger"
	"github.com/gin-gonic/gin"
)

func RequestLogger(logger *logger.Logger) gin.HandlerFunc {
	requestLogger := logger.With().Str("service", "REQS").Logger()
	return func(c *gin.Context) {
		c.Set("logger", requestLogger)
		c.Next()
	}
}
