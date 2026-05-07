package user

import (
	"context"

	"gorm.io/gorm"
)

// Repository is User repository that handles database requests for Services
type Repository interface {
	// Create creates adds new user to database.
	//
	// tx parameter should be set nil outside of transactions
	//
	// returns gorm error on internal errors.
	Create(ctx context.Context, tx *gorm.DB, user User) (User, error)
}

type repository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

// WithTx is used for creating transactions
// for making sure when a failure happens all database actions will be cancelled.
//
// While not using transactions tx parameter should be nil
func (r *repository) WithTx(tx *gorm.DB) *gorm.DB {
	if tx != nil {
		return tx
	}
	return r.db
}

// Create creates adds new user to database.
//
// tx parameter should be set nil outside of transactions
//
// returns gorm error on internal errors.
func (r *repository) Create(ctx context.Context, tx *gorm.DB, user User) (User, error) {
	if err := r.WithTx(tx).WithContext(ctx).Create(&user).Error; err != nil {
		return User{}, err
	}
	return user, nil
}
