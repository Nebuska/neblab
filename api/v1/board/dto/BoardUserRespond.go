package dto

import "github.com/Nebuska/task-tracker/internal/boardUser"

type BoardUserRespond struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Role string `json:"role"`
}

func NewBoardUserRespond(boardUser boardUser.BoardUser) BoardUserRespond {
	return BoardUserRespond{
		ID:   boardUser.UserID,
		Role: boardUser.Role,
	}
}

func NewBoardUsersRespond(boardUsers []boardUser.BoardUser) []BoardUserRespond {
	boardUsersRespond := make([]BoardUserRespond, len(boardUsers))
	for i := range boardUsersRespond {
		boardUsersRespond[i] = NewBoardUserRespond(boardUsers[i])
	}
	return boardUsersRespond
}
