package task

import (
	"net/http"
	"strconv"

	"github.com/Nebuska/task-tracker/api/v1/task/dto"
	"github.com/Nebuska/task-tracker/internal/task"
	"github.com/Nebuska/task-tracker/pkg/appError"
	"github.com/Nebuska/task-tracker/pkg/appError/errorCodes"
	"github.com/Nebuska/task-tracker/pkg/jwtAuth"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service task.Service
}

func NewTaskHandler(service task.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) GetTasks(context *gin.Context) {
	claims := context.MustGet("claims").(*jwtAuth.UserClaims)
	boardId, err := strconv.ParseUint(context.Param("id"), 10, 64)
	if err != nil {
		_ = context.Error(appError.New(errorCodes.BadRequest, "TaskHandler", err.Error()))
		return
	}
	tasksModel, err := h.service.GetTasksByBoardUsingUser(claims.UserId, uint(boardId))
	if err != nil {
		_ = context.Error(err)
		return
	}
	context.JSON(http.StatusOK, dto.NewTasksRespond(tasksModel))
}

func (h *Handler) GetTask(context *gin.Context) {
	claims := context.MustGet("claims").(*jwtAuth.UserClaims)
	taskId, err := strconv.ParseUint(context.Param("id"), 10, 64)
	if err != nil {
		_ = context.Error(appError.New(errorCodes.BadRequest, "TaskHandler", err.Error()))
		return
	}
	taskModel, err := h.service.GetTaskByIdUsingUser(claims.UserId, uint(taskId))
	if err != nil {
		_ = context.Error(err)
		return
	}
	context.JSON(http.StatusOK, dto.NewTaskRespond(taskModel))
}

func (h *Handler) CreateTask(context *gin.Context, requestDto dto.CreateTaskRequest) {
	claims := context.MustGet("claims").(*jwtAuth.UserClaims)
	newTaskModel, err := h.service.CreateTaskUsingUser(claims.UserId, requestDto.ToModel())
	if err != nil {
		_ = context.Error(err)
		return
	}

	context.JSON(http.StatusCreated, dto.NewTaskRespond(newTaskModel))
}
