package v1

import (
	"github.com/Nebuska/neblab/account/api/middlewares"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

func RegisterRoutes(engine *gin.Engine, authHandler *AuthHandler) {
	router := engine.Group("/api/v1")

	router.POST("/jwt", middlewares.WithBody(authHandler.GetJwt))
	router.POST("/register", middlewares.WithBody(authHandler.Register))
}

var Module = fx.Options(
	fx.Provide(NewAuthHandler),
	fx.Invoke(RegisterRoutes),
)
