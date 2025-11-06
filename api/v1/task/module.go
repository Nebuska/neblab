package task

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

func RegisterRoutes(engine *gin.Engine, handler *Handler) {
	router := engine.Group("/api/v1/tasks")

	router.GET("", handler.GetTasks)
	router.POST("", handler.CreateTask)
}

var Module = fx.Options(
	fx.Provide(NewTaskHandler),
	fx.Invoke(RegisterRoutes),
)
