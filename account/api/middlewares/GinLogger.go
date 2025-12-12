package middlewares

import (
	"time"

	"github.com/Nebuska/neblab/shared/logger"
	"github.com/gin-gonic/gin"
)

func GinLogger(logger *logger.Logger) gin.HandlerFunc {
	ginLogger := logger.With().Str("service", " GIN").Logger()

	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		duration := time.Since(start)
		ginLogger.Info().
			Str("method", c.Request.Method).
			Str("path", c.Request.URL.Path).
			Int("status", c.Writer.Status()).
			Str("latency", duration.String()).
			Send()
	}
}
