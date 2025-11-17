package board

import (
	"net/http"
	"strconv"

	"github.com/Nebuska/task-tracker/api/v1/board/dto"
	"github.com/Nebuska/task-tracker/internal/board"
	"github.com/Nebuska/task-tracker/pkg/appError"
	"github.com/Nebuska/task-tracker/pkg/appError/errorCodes"
	"github.com/Nebuska/task-tracker/pkg/jwtAuth"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service board.Service
}

func NewBoardHandler(service board.Service) *Handler {
	return &Handler{service: service}

}

func (h Handler) GetBoard(context *gin.Context) {
	claims := context.MustGet("claims").(*jwtAuth.UserClaims)
	boardId, err := strconv.ParseUint(context.Param("id"), 10, 64)
	if err != nil {
		_ = context.Error(appError.New(errorCodes.BadRequest, "BoardHandler", err.Error()))
	}
	boardModel, err := h.service.GetBoardUsingUser(claims.UserId, uint(boardId))
	if err != nil {
		_ = context.Error(err)
		return
	}
	context.JSON(http.StatusOK, dto.NewBoardRespond(boardModel))
}

func (h Handler) CreateBoard(context *gin.Context, requestDTO dto.CreateBoardRequest) {
	claims := context.MustGet("claims").(*jwtAuth.UserClaims)
	newBoardModel, err := h.service.CreateBoardUsingUser(claims.UserId, requestDTO.ToModel())
	if err != nil {
		_ = context.Error(err)
		return
	}
	context.JSON(http.StatusOK, dto.NewBoardRespond(newBoardModel))
}

func (h Handler) GetUserBoards(context *gin.Context) {
	claims := context.MustGet("claims").(*jwtAuth.UserClaims)
	boardsModel, err := h.service.GetBoardsUsingUser(claims.UserId)
	if err != nil {
		_ = context.Error(err)
		return
	}
	context.JSON(http.StatusOK, dto.NewBoardsRespond(boardsModel))
}

func (h Handler) DeleteBoard(context *gin.Context) {
	claims := context.MustGet("claims").(*jwtAuth.UserClaims)
	boardId, err := strconv.ParseUint(context.Param("id"), 10, 64)
	if err != nil {
		_ = context.Error(appError.New(errorCodes.BadRequest, "BoardHandler", err.Error()))
	}
	err = h.service.DeleteBoardUsingUser(claims.UserId, uint(boardId))
	if err != nil {
		_ = context.Error(err)
		return
	}
	context.JSON(http.StatusOK, "Deleted board")
}
