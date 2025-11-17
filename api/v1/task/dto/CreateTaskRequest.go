package dto

import "github.com/Nebuska/task-tracker/internal/task"

type CreateTaskRequest struct {
	Name        string `json:"name" binding:"required,min=3,max=50"`
	Description string `json:"description" binding:"omitempty,min=5"`
	Status      string `json:"status" binding:"oneof=Planning InProgress Completed"`
	BoardID     uint   `json:"board_id" binding:"required"`
}

func (receiver CreateTaskRequest) ToModel() task.Task {
	return task.Task{
		Name:        receiver.Name,
		Description: receiver.Description,
		Status:      receiver.Status,
		BoardID:     receiver.BoardID,
	}
}
