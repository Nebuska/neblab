package Auth

import (
	"task-tracker/internal/User"

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
	//TODO implement me
	user := UserCredentials{
		Username: username,
		Password: password,
		User: User.User{
			Email: email,
		},
	}

	return user, repo.db.Create(&user).Error
}

func (repo authRepository) GetCredentials(username string) (UserCredentials, error) {
	var user UserCredentials
	repo.db.Where("username = ?", username).First(&user)
	return user, repo.db.Error
}
