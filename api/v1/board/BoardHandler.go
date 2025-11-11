package board

import (
	"net/http"
	"strconv"
	"task-tracker/internal/Board"
	"task-tracker/pkg/appError"
	"task-tracker/pkg/appError/errorCodes"
	"task-tracker/pkg/jwtAuth"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service Board.Service
}

func NewBoardHandler(service Board.Service) *Handler {
	return &Handler{service: service}

}

func (h Handler) GetBoard(context *gin.Context) {
	claims := context.MustGet("claims").(*jwtAuth.UserClaims)
	boardId, err := strconv.ParseUint(context.Param("id"), 10, 64)
	if err != nil {
		_ = context.Error(appError.New(errorCodes.BadRequest, "BoardHandler", err.Error()))
	}
	board, err := h.service.GetBoardUsingUser(claims.UserId, uint(boardId))
	if err != nil {
		_ = context.Error(err)
		return
	}
	context.JSON(http.StatusOK, board)
}

func (h Handler) CreateBoard(context *gin.Context) {
	claims := context.MustGet("claims").(*jwtAuth.UserClaims)
	var board Board.Board
	if err := context.ShouldBindJSON(&board); err != nil {
		_ = context.Error(appError.New(errorCodes.BadRequest, "BoardHandler", err.Error()))
		return
	}
	board, err := h.service.CreateBoardUsingUser(claims.UserId, board)
	if err != nil {
		_ = context.Error(err)
		return
	}
	context.JSON(http.StatusOK, board)
}

func (h Handler) GetUserBoards(context *gin.Context) {
	claims := context.MustGet("claims").(*jwtAuth.UserClaims)
	boards, err := h.service.GetBoardsUsingUser(claims.UserId)
	if err != nil {
		_ = context.Error(err)
		return
	}
	context.JSON(http.StatusOK, boards)
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
