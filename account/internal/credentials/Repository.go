package credentials

import (
	"context"
	"errors"

	"github.com/Nebuska/neblab/account/internal/dto"
	"gorm.io/gorm"
)

// Repository is Credentials repository that handles database requests for Services
type Repository interface {
	// Create creates adds new credentials to database.
	//
	// tx parameter should be set nil outside of transactions
	//
	// It returns the credentials on success with no error.
	// returns dto.ErrUserAlreadyExist error if username already has credentials in database.
	// returns gorm error on internal errors.
	Create(ctx context.Context, tx *gorm.DB, credentials Credentials) (Credentials, error)
	// Find searches username on credentials and brings the matching credential
	//
	// tx parameter should be set nil outside of transactions
	//
	// It returns the matching credential on success with no error
	// returns dto.ErrUserNotFound error if username is not inside the database
	// returns gorm error on internal errors
	Find(ctx context.Context, tx *gorm.DB, username string) (Credentials, error)
	// Deprecated: No longer used since Create handles checks with an error
	CheckEmail(ctx context.Context, tx *gorm.DB, email string) (Credentials, bool, error)
}

type repository struct {
	db *gorm.DB
}

func NewCredentialsRepository(db *gorm.DB) Repository {
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

// Create creates adds new credentials to database.
//
// tx parameter should be set nil outside of transactions
//
// It returns the credentials on success with no error.
// returns dto.ErrUserAlreadyExist error if username already has credentials in database.
// returns gorm error on internal errors.
func (r *repository) Create(ctx context.Context, tx *gorm.DB, credentials Credentials) (Credentials, error) {
	if err := r.WithTx(tx).WithContext(ctx).Create(&credentials).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return Credentials{}, dto.ErrUserAlreadyExist
		}
		return Credentials{}, err
	}
	return credentials, nil
}

// Find searches username on credentials and brings the matching credential
//
// tx parameter should be set nil outside of transactions
//
// It returns the matching credential on success with no error
// returns dto.ErrUserNotFound error if username is not inside the database
// returns gorm error on internal errors
func (r *repository) Find(ctx context.Context, tx *gorm.DB, username string) (Credentials, error) {
	var credentials Credentials
	if err := r.WithTx(tx).WithContext(ctx).Where("username = ?", username).Take(&credentials).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return Credentials{}, dto.ErrUserNotFound
		}
		return Credentials{}, err
	}
	return credentials, nil
}

// Deprecated: No longer used since Create handles checks with an error
// it should be changed with username if wanted to used again
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
