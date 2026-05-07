package v1

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewAuthHandler),

	fx.Invoke(RegisterRoutes),
)

func RegisterRoutes(engine *gin.Engine,
	authHandler *AuthHandler,
) {
	api := engine.Group("/v1")

	authHandler.RegisterRoute(api)
}
