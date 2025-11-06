package Task

import (
	"gorm.io/gorm"
)

type TaskRepository struct {
	DB *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *TaskRepository {
	return &TaskRepository{DB: db}
}

func (r *TaskRepository) GetAll() ([]Task, error) {
	var tasks []Task
	err := r.DB.Find(&tasks).Error
	return tasks, err
}

func (r *TaskRepository) GetByID(id int) (Task, error) {
	var task Task
	err := r.DB.First(&task, id).Error
	return task, err
}

func (r *TaskRepository) Create(task Task) (Task, error) {
	err := r.DB.Create(&task).Error
	return task, err
}

func (r *TaskRepository) Update(task Task) (Task, error) {
	err := r.DB.Save(&task).Error
	return task, err
}

func (r *TaskRepository) Delete(task Task) error {
	err := r.DB.Delete(&task).Error
	return err
}
