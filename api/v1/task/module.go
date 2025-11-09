package task

import (
	"task-tracker/api/middlewares"
	"task-tracker/pkg/jwtAuth"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

func RegisterRoutes(engine *gin.Engine, handler *Handler, manager *jwtAuth.JWTManager) {
	router := engine.Group("/api/v1/tasks")
	router.Use(middlewares.AuthMiddleware(manager))
	{
		router.GET("/fromBoard/:id", handler.GetTasks)
		router.GET("/:id", handler.GetTask)
		router.POST("", handler.CreateTask)
	}
}

var Module = fx.Options(
	fx.Provide(NewTaskHandler),
	fx.Invoke(RegisterRoutes),
)
