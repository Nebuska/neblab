package task

import (
	"github.com/Nebuska/task-tracker/pkg/appError"
	"github.com/Nebuska/task-tracker/pkg/appError/errorCodes"
)

type Service interface {
	CreateTaskUsingUser(userId uint, task Task) (Task, error)
	GetTasksByBoardUsingUser(userId uint, boardId uint) ([]Task, error)
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

func (s *taskService) GetTasksByBoardUsingUser(userId uint, boardId uint) ([]Task, error) {
	hasAccess, err := s.repo.UserHasAccessToBoard(userId, boardId)
	if err != nil {
		return nil, err
	}
	if !hasAccess {
		return nil, appError.New(errorCodes.Forbidden,
			"TaskService",
			"user does not have access to task's board")
	}

	return s.repo.GetTasksByBoard(boardId)

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
