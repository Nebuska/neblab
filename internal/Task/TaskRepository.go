package Task

import (
	"task-tracker/internal/common"

	"gorm.io/gorm"
)

type Repository interface {
	CreateTask(task Task) (Task, error)
	GetTasksByBoard(boardId uint) ([]Task, error)
	GetTaskById(taskId uint) (Task, error)
	UserHasAccessToBoard(userID, boardID uint) (bool, error)
}

type taskRepository struct {
	common.Repository[Task]
}

func NewTaskRepository(db *gorm.DB) Repository {
	return &taskRepository{
		Repository: common.Repository[Task]{
			DB: db,
		},
	}
}

func (repo *taskRepository) UserHasAccessToBoard(userID, boardID uint) (bool, error) {
	var count int64
	err := repo.DB.Table("board_users").
		Where("boardId = ? AND user_id = ?", boardID, userID).
		Count(&count).Error
	return count > 0, err
}

func (repo *taskRepository) CreateTask(task Task) (Task, error) {
	err := repo.DB.Create(&task).Error
	return task, err
}

func (repo *taskRepository) GetTasksByBoard(boardId uint) ([]Task, error) {
	var tasks []Task
	err := repo.DB.
		Where("board_id = ?", boardId).Find(&tasks).Error
	return tasks, err
}

func (repo *taskRepository) GetTaskById(taskId uint) (Task, error) {
	task := Task{Model: gorm.Model{ID: taskId}}
	err := repo.DB.
		Where("tasks.id = ?", taskId).First(&task).Error
	return task, err
}
