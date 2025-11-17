package dto

import "github.com/Nebuska/task-tracker/internal/board"

type CreateBoardRequest struct {
	Name string `json:"name" binding:"required,min=3,max=30"`
}

func (receiver CreateBoardRequest) ToModel() board.Board {
	return board.Board{
		Name: receiver.Name,
	}
}
