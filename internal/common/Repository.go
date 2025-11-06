package common

import "gorm.io/gorm"

type Repository[T any] struct {
	DB *gorm.DB
}

func (r *Repository[T]) GetAll() ([]T, error) {
	var t []T
	err := r.DB.Find(&t).Error
	return t, err
}

func (r *Repository[T]) GetByID(id uint) (T, error) {
	var t T
	err := r.DB.First(&t, id).Error
	return t, err
}

func (r *Repository[T]) Create(t T) (T, error) {
	err := r.DB.Create(&t).Error
	return t, err
}

func (r *Repository[T]) Update(t T) (T, error) {
	err := r.DB.Save(&t).Error
	return t, err
}

func (r *Repository[T]) Delete(t T) error {
	err := r.DB.Delete(&t).Error
	return err
}
