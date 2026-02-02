package api

import (
	v1 "github.com/Nebuska/neblab/account/api/v1"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

func RegisterRoutes(engine *gin.Engine) {
	router := engine.Group("/api/v1")

	router.GET("/health", health)
}

var Module = fx.Options(
	fx.Invoke(RegisterRoutes),
	v1.Module,
)
