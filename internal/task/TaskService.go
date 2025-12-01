package task

import (
	"github.com/Nebuska/task-tracker/pkg/appError"
	"github.com/Nebuska/task-tracker/pkg/appError/errorCodes"
)

type Service interface {
	CreateTaskUsingUser(userId uint, task Task) (Task, error)
	GetTasksByFilterUsingUser(userId uint, filter Filter) ([]Task, error)
	GetTaskByIdUsingUser(userID uint, taskId uint) (Task, error)
}

type taskService struct {
	repo Repository
}

func NewTaskService(repo Repository) Service {
	return &taskService{repo: repo}
}

func (s *taskService) CreateTaskUsingUser(userId uint, task Task) (Task, error) {
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

func (s *taskService) GetTasksByFilterUsingUser(userId uint, filter Filter) ([]Task, error) {
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

func (s *taskService) GetTaskByIdUsingUser(userId uint, taskId uint) (Task, error) {
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

func (s *taskService) CreateTask(task Task) (Task, error) {
	//TODO implement me
	panic("implement me")

}

func (s *taskService) GetTasksByBoard(boardId uint) ([]Task, error) {
	//TODO implement me
	panic("implement me")
}

func (s *taskService) GetTaskById(taskId uint) (Task, error) {
	//TODO implement me
	panic("implement me")
}
