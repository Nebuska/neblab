package task

import (
	"github.com/Nebuska/neblab/shared/appError"
	"github.com/Nebuska/neblab/shared/appError/errorCodes"

	"gorm.io/gorm"
)

type Repository interface {
	CreateTask(task Task) (Task, error)
	GetTasksByFilter(filter Filter) ([]Task, error)
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

func (repo *taskRepository) GetTasksByFilter(filter Filter) ([]Task, error) {
	var tasks []Task
	gormQuery := repo.DB.Table("tasks")
	if len(filter.BoardId) > 0 {
		gormQuery = gormQuery.Where("board_id IN (?)", filter.BoardId)
	}
	if len(filter.Status) > 0 {
		gormQuery = gormQuery.Where("status IN (?)", filter.Status)
	}
	if filter.Search != "" {
		gormQuery = gormQuery.Where("title LIKE ?", "%"+filter.Search+"%")
	}
	gormQuery = gormQuery.Joins("JOIN boards ON tasks.board_id = boards.id").
		Joins("JOIN board_users ON board_users.board_id = boards.id").
		Where("board_users.user_id = ?", filter.AccessibleByUserId).
		Order(filter.OrderBy())

	if filter.Before != 0 || filter.After != 0 {
		//todo : advanced query options with before and after
		return nil, appError.New(errorCodes.BadRequest, "TaskRepo", "Not Implemented")
	}

	err := gormQuery.Offset(filter.PageSize * (filter.PageNumber - 1)).
		Limit(filter.PageSize).Find(&tasks).Error

	return tasks, appError.FromGormError(err)
}

func (repo *taskRepository) GetTaskById(taskId uint) (Task, error) {
	task := Task{Model: gorm.Model{ID: taskId}}
	err := repo.DB.
		Where("tasks.id = ?", taskId).First(&task).Error
	return task, appError.FromGormError(err)
}
