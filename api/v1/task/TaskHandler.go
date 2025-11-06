package task

import (
	"net/http"
	"task-tracker/internal/Task"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service Task.Service
}

func NewTaskHandler(service Task.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) GetTasks(context *gin.Context) {
	tasks, err := h.service.GetAll()
	if err != nil {
		context.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	context.JSON(http.StatusOK, tasks)
}

func (h *Handler) CreateTask(context *gin.Context) {
	var task Task.Task
	if err := context.ShouldBindJSON(&task); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newTask, err := h.service.Create(task)
	if err != nil {
		//todo : give better errors
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, newTask)
}
