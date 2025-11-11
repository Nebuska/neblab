package aTask

import (
	"github.com/Nebuska/task-tracker/pkg/appError"

	"gorm.io/gorm"
)

type Repository interface {
	CreateTask(task Task) (Task, error)
	GetTasksByBoard(boardId uint) ([]Task, error)
	GetTaskById(taskId uint) (Task, error)
	UserHasAccessToBoard(userID, boardID uint) (bool, error)
}

type taskRepository struct {
	DB *gorm.DB
}

func NewTaskRepository(db *gorm.DB) Repository {
	return &taskRepository{
		DB: db,
	}
}

func (repo *taskRepository) UserHasAccessToBoard(userID, boardID uint) (bool, error) {
	var count int64
	err := repo.DB.Table("board_users").
		Where("board_id = ? AND user_id = ?", boardID, userID).
		Count(&count).Error
	return count > 0, appError.FromGormError(err)
}

func (repo *taskRepository) CreateTask(task Task) (Task, error) {
	err := repo.DB.Create(&task).Error
	return task, appError.FromGormError(err)
}

func (repo *taskRepository) GetTasksByBoard(boardId uint) ([]Task, error) {
	var tasks []Task
	err := repo.DB.
		Where("board_id = ?", boardId).Find(&tasks).Error
	return tasks, appError.FromGormError(err)
}

func (repo *taskRepository) GetTaskById(taskId uint) (Task, error) {
	task := Task{Model: gorm.Model{ID: taskId}}
	err := repo.DB.
		Where("tasks.id = ?", taskId).First(&task).Error
	return task, appError.FromGormError(err)
}
