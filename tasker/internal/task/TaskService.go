package task

import (
	"github.com/Nebuska/neblab/shared/appError"
	"github.com/Nebuska/neblab/shared/appError/errorCodes"
)

type Service interface {
	CreateTask(userId uint, task Task) (Task, error)
	GetTasksByFilter(userId uint, filter Filter) ([]Task, error)
	GetTaskById(userID uint, taskId uint) (Task, error)
}

type taskService struct {
	repo Repository
}

func NewTaskService(repo Repository) Service {
	return &taskService{repo: repo}
}

func (s *taskService) CreateTask(userId uint, task Task) (Task, error) {
	if task.BoardID == 0 {
		return Task{}, appError.New(errorCodes.BadRequest,
			"TaskRepo",
			"board id is required")
	}
	task.ID = 0
	hasAccess, err := s.repo.UserHasAccessToBoard(userId, task.BoardID)
	if err != nil {
		return Task{}, err
	}
	if !hasAccess {
		return Task{}, appError.New(errorCodes.Forbidden,
			"TaskService",
			"user does not have access to task's board")
	}
	return s.repo.CreateTask(task)
}

func (s *taskService) GetTasksByFilter(userId uint, filter Filter) ([]Task, error) {
	filter.AccessibleByUserId = userId
	if filter.PageSize == 0 {
		filter.PageSize = 20
	}
	if filter.SortBy == "" {
		filter.SortBy = "created_at"
		filter.ReversedSort = true
	}

	return s.repo.GetTasksByFilter(filter)

}

func (s *taskService) GetTaskById(userId uint, taskId uint) (Task, error) {
	task, err := s.repo.GetTaskById(taskId)
	if err != nil {
		return Task{}, err
	}
	hasAccess, err := s.repo.UserHasAccessToBoard(userId, task.BoardID)
	if err != nil {
		return Task{}, err
	}
	if !hasAccess {
		return Task{}, appError.New(errorCodes.Forbidden,
			"TaskService",
			"user does not have access to task's board")
	}
	return task, err
}
