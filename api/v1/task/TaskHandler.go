package task

import (
	"task-tracker/internal/Task"

	"github.com/gin-gonic/gin"
)

type TaskHandler struct {
	service *Task.TaskService
}

func NewTaskHandler(service *Task.TaskService) *TaskHandler {
	return &TaskHandler{service: service}
}

func (h *TaskHandler) GetTasks(context *gin.Context) {
	context.JSON(200, gin.H{"test": "test"})
	//todo : not implemented
}
