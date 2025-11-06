package common

type CRUD[T any] interface {
	GetAll() ([]T, error)
	GetByID(id uint) (T, error)
	Create(t T) (T, error)
	Update(t T) (T, error)
	Delete(t T) error
}
