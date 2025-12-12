package credentials

import (
	"context"
	"errors"

	"gorm.io/gorm"
)

type Repository interface {
	Create(ctx context.Context, tx *gorm.DB, credentials Credentials) (Credentials, error)
	Find(ctx context.Context, tx *gorm.DB, username string) (Credentials, error)
	CheckEmail(ctx context.Context, tx *gorm.DB, email string) (Credentials, bool, error)
}

type repository struct {
	db *gorm.DB
}

func NewCredentialsRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) WithTx(tx *gorm.DB) *gorm.DB {
	if tx != nil {
		return tx
	}
	return r.db
}

func (r *repository) Create(ctx context.Context, tx *gorm.DB, credentials Credentials) (Credentials, error) {
	if err := r.WithTx(tx).WithContext(ctx).Create(&credentials).Error; err != nil {
		return Credentials{}, err
	}
	return credentials, nil
}

func (r *repository) Find(ctx context.Context, tx *gorm.DB, username string) (Credentials, error) {
	var credentials Credentials
	if err := r.WithTx(tx).WithContext(ctx).Where("username = ?", username).Take(&credentials).Error; err != nil {
		return Credentials{}, err
	}
	return credentials, nil
}

func (r *repository) CheckEmail(ctx context.Context, tx *gorm.DB, email string) (Credentials, bool, error) {
	var credentials Credentials
	err := r.WithTx(tx).WithContext(ctx).Where("email = ?", email).Take(credentials).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return Credentials{}, false, nil
		}
		return Credentials{}, false, err
	}
	return credentials, true, nil
}
