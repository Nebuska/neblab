package task

import (
	"net/http"
	"strconv"
	"task-tracker/internal/Task"
	"task-tracker/pkg/jwtAuth"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service Task.Service
}

func NewTaskHandler(service Task.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) GetTasks(context *gin.Context) {
	claims := context.MustGet("claims").(*jwtAuth.UserClaims)
	boardId, _ := strconv.ParseUint(context.Param("Id"), 10, 64)
	tasks, err := h.service.GetTasksByBoardUsingUser(claims.UserId, uint(boardId))
	if err != nil {
		context.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	context.JSON(http.StatusOK, tasks)
}

func (h *Handler) GetTask(context *gin.Context) {
	claims := context.MustGet("claims").(*jwtAuth.UserClaims)
	taskId, _ := strconv.ParseUint(context.Param("Id"), 10, 64)
	tasks, err := h.service.GetTaskByIdUsingUser(claims.UserId, uint(taskId))
	if err != nil {
		context.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	context.JSON(http.StatusOK, tasks)
}

func (h *Handler) CreateTask(context *gin.Context) {
	claims := context.MustGet("claims").(*jwtAuth.UserClaims)
	var task Task.Task
	if err := context.ShouldBindJSON(&task); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newTask, err := h.service.CreateTaskUsingUser(claims.UserId, task)
	if err != nil {
		//todo : give better errors
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, newTask)
}
