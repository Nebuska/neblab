package user

import (
	"github.com/Nebuska/neblab/account/internal/dto"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	FirstName string
	LastName  string

	Email string
}

func FromRegisterDTO(dto dto.RegisterData) User {
	return User{
		FirstName: "",
		LastName:  "",
		Email:     dto.Email,
	}
}
