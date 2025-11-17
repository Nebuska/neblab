package dto

import "github.com/Nebuska/task-tracker/internal/task"

type TaskRespond struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

func NewTaskRespond(task task.Task) TaskRespond {
	return TaskRespond{
		ID:          task.ID,
		Name:        task.Name,
		Description: task.Description,
		Status:      task.Status,
	}
}

func NewTasksRespond(tasks []task.Task) []TaskRespond {
	tasksRespond := make([]TaskRespond, len(tasks))
	for i := range tasks {
		tasksRespond[i] = NewTaskRespond(tasks[i])
	}
	return tasksRespond
}
