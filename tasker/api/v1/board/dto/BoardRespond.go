package dto

import (
	taskDTO "github.com/Nebuska/neblab/tasker/api/v1/task/dto"
	"github.com/Nebuska/neblab/tasker/internal/board"
)

type BoardRespond struct {
	ID         uint                  `json:"id"`
	Name       string                `json:"name"`
	Tasks      []taskDTO.TaskRespond `json:"tasks"`
	BoardUsers []BoardUserRespond    `json:"board_users"`
}

func NewBoardRespond(board board.Board) BoardRespond {
	return BoardRespond{
		ID:         board.ID,
		Name:       board.Name,
		Tasks:      taskDTO.NewTasksRespond(board.Tasks),
		BoardUsers: NewBoardUsersRespond(board.BoardUsers),
	}
}

func NewBoardsRespond(boards []board.Board) []BoardRespond {
	boardsRespond := make([]BoardRespond, len(boards))
	for i := range boards {
		boardsRespond[i] = NewBoardRespond(boards[i])
	}
	return boardsRespond
}
