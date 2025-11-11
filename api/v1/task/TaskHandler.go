package task

import (
	"github.com/Nebuska/task-tracker/internal/aTask"
	"github.com/Nebuska/task-tracker/pkg/appError"
	"github.com/Nebuska/task-tracker/pkg/appError/errorCodes"
	"github.com/Nebuska/task-tracker/pkg/jwtAuth"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service aTask.Service
}

func NewTaskHandler(service aTask.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) GetTasks(context *gin.Context) {
	claims := context.MustGet("claims").(*jwtAuth.UserClaims)
	boardId, err := strconv.ParseUint(context.Param("id"), 10, 64)
	if err != nil {
		_ = context.Error(appError.New(errorCodes.BadRequest, "TaskHandler", err.Error()))
		return
	}
	tasks, err := h.service.GetTasksByBoardUsingUser(claims.UserId, uint(boardId))
	if err != nil {
		_ = context.Error(err)
		return
	}
	context.JSON(http.StatusOK, tasks)
}

func (h *Handler) GetTask(context *gin.Context) {
	claims := context.MustGet("claims").(*jwtAuth.UserClaims)
	taskId, err := strconv.ParseUint(context.Param("id"), 10, 64)
	if err != nil {
		_ = context.Error(appError.New(errorCodes.BadRequest, "TaskHandler", err.Error()))
		return
	}
	tasks, err := h.service.GetTaskByIdUsingUser(claims.UserId, uint(taskId))
	if err != nil {
		_ = context.Error(err)
		return
	}
	context.JSON(http.StatusOK, tasks)
}

func (h *Handler) CreateTask(context *gin.Context) {
	claims := context.MustGet("claims").(*jwtAuth.UserClaims)
	var task aTask.Task
	if err := context.ShouldBindJSON(&task); err != nil {
		_ = context.Error(appError.New(errorCodes.BadRequest, "TaskHandler", err.Error()))
		return
	}

	newTask, err := h.service.CreateTaskUsingUser(claims.UserId, task)
	if err != nil {
		_ = context.Error(err)
		return
	}

	context.JSON(http.StatusCreated, newTask)
}
