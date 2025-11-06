package Task

import (
	"errors"
)

type TaskService struct {
	repo *TaskRepository
}

func NewTaskService(repo *TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) Create(task Task) (Task, error) {
	if task.Name == "" {
		return Task{}, errors.New("task name is required")
	}
	if task.BoardID == 0 {
		return Task{}, errors.New("task board id is required")
	}
	if task.Status == "" {
		task.Status = "todo"
	}
	return s.repo.Create(task)
}

func (s *TaskService) Update(task Task) (Task, error) {
	if task.ID == 0 {
		return Task{}, errors.New("task id is required")
	}
	return s.repo.Update(task)
}

func (s *TaskService) Delete(task Task) error {
	if task.ID == 0 {
		return errors.New("task id is required")
	}
	return s.repo.Delete(task)
}
