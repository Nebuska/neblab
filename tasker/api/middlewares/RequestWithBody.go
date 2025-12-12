package middlewares

import (
	"github.com/Nebuska/neblab/shared/appError"
	"github.com/Nebuska/neblab/shared/appError/errorCodes"
	"github.com/gin-gonic/gin"
)

func WithBody[T any](handler func(ctx *gin.Context, t T)) gin.HandlerFunc {
	return func(c *gin.Context) {
		var t T
		err := c.ShouldBindBodyWithJSON(&t)
		if err != nil {
			_ = c.Error(appError.New(errorCodes.BadRequest, "RequestWithBody", err.Error()))
			return
		}
		handler(c, t)
	}
}
