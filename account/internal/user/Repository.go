package user

import (
	"context"

	"gorm.io/gorm"
)

type Repository interface {
	Create(ctx context.Context, tx *gorm.DB, user User) (User, error)
}

type repository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) WithTx(tx *gorm.DB) *gorm.DB {
	if tx != nil {
		return tx
	}
	return r.db
}

func (r *repository) Create(ctx context.Context, tx *gorm.DB, user User) (User, error) {
	if err := r.WithTx(tx).WithContext(ctx).Create(&user).Error; err != nil {
		return User{}, err
	}
	return user, nil
}
