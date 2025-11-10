package board

import (
	"task-tracker/api/middlewares"
	"task-tracker/pkg/jwtAuth"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

func RegisterRoutes(engine *gin.Engine, handler *Handler, manager *jwtAuth.JWTManager) {
	router := engine.Group("/api/v1/boards")
	router.Use(middlewares.AuthMiddleware(manager))
	{
		router.GET("", handler.GetUserBoards)
		router.GET("/:id", handler.GetBoard)
		router.POST("", handler.CreateBoard)

		router.DELETE(":id", handler.DeleteBoard)
	}
}

var Module = fx.Options(
	fx.Provide(NewBoardHandler),
	fx.Invoke(RegisterRoutes),
)
