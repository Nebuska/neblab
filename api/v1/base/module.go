package base

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

func RegisterRoutes(engine *gin.Engine, authHandler *AuthHandler) {
	router := engine.Group("/api/v1")

	router.POST("/login", authHandler.Login)
	router.POST("/register", authHandler.Register)
}

var Module = fx.Options(
	fx.Provide(NewAuthHandler),
	fx.Invoke(RegisterRoutes),
)
