package auth

import (
	"github.com/Nebuska/task-tracker/internal/user"
	"github.com/Nebuska/task-tracker/pkg/appError"

	"gorm.io/gorm"
)

type Repository interface {
	Register(username, email, password string) (UserCredentials, error)
	GetCredentials(username string) (UserCredentials, error)
}

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) Repository {
	return &authRepository{db: db}
}

func (repo authRepository) Register(username, email, password string) (UserCredentials, error) {
	user := UserCredentials{
		Username: username,
		Password: password,
		User: user.User{
			Email: email,
		},
	}
	err := repo.db.Create(&user).Error
	return user, appError.FromGormError(err)
}

func (repo authRepository) GetCredentials(username string) (UserCredentials, error) {
	var user UserCredentials
	err := repo.db.Where("username = ?", username).First(&user).Error
	return user, appError.FromGormError(err)
}
