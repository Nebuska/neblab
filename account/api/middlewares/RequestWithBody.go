package middlewares

import (
	"github.com/gin-gonic/gin"
)

func WithBody[T any](handler func(ctx *gin.Context, body T)) gin.HandlerFunc {
	return func(c *gin.Context) {
		var body T
		if err := c.ShouldBind(&body); err != nil {
			//todo : identify the error
			_ = c.Error(err)
			return
		}
		handler(c, body)
	}
}
