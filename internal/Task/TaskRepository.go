package Task

import (
	"task-tracker/internal/common"

	"gorm.io/gorm"
)

type Repository interface {
	common.CRUD[Task]
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
