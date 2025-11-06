package Task

import (
	"errors"
	"task-tracker/internal/common"
)

type Service interface {
	common.CRUD[Task]
}

type taskService struct {
	repo Repository
}

func NewTaskService(repo Repository) Service {
	return &taskService{repo: repo}
}

func (s *taskService) GetAll() ([]Task, error) {
	return s.repo.GetAll()
}

func (s *taskService) GetByID(id uint) (Task, error) {
	return s.repo.GetByID(id)
}

func (s *taskService) Create(task Task) (Task, error) {
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

func (s *taskService) Update(task Task) (Task, error) {
	if task.ID == 0 {
		return Task{}, errors.New("task id is required")
	}
	return s.repo.Update(task)
}

func (s *taskService) Delete(task Task) error {
	if task.ID == 0 {
		return errors.New("task id is required")
	}
	return s.repo.Delete(task)
}
