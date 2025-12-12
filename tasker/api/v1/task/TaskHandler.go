package task

import (
	"net/http"
	"strconv"

	"github.com/Nebuska/neblab/shared/appError"
	"github.com/Nebuska/neblab/shared/appError/errorCodes"
	"github.com/Nebuska/neblab/shared/jwtAuth"
	"github.com/Nebuska/neblab/tasker/api/v1/task/dto"
	"github.com/Nebuska/neblab/tasker/internal/task"

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

	var query dto.TaskQuery
	err := context.ShouldBindQuery(&query)
	if err != nil {
		_ = context.Error(appError.New(errorCodes.BadRequest, "TaskHandler", err.Error()))
		return
	}
	tasksModel, err := h.service.GetTasksByFilter(claims.UserId, query.ToFilter())
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
	taskModel, err := h.service.GetTaskById(claims.UserId, uint(taskId))
	if err != nil {
		_ = context.Error(err)
		return
	}
	context.JSON(http.StatusOK, dto.NewTaskRespond(taskModel))
}

func (h *Handler) CreateTask(context *gin.Context, requestDto dto.CreateTaskRequest) {
	claims := context.MustGet("claims").(*jwtAuth.UserClaims)
	newTaskModel, err := h.service.CreateTask(claims.UserId, requestDto.ToModel())
	if err != nil {
		_ = context.Error(err)
		return
	}

	context.JSON(http.StatusCreated, dto.NewTaskRespond(newTaskModel))
}
