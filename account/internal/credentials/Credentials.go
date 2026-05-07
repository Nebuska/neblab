package credentials

import (
	"github.com/Nebuska/neblab/account/internal/dto"
	"github.com/Nebuska/neblab/account/internal/user"

	"gorm.io/gorm"
)

type Credentials struct {
	gorm.Model
	Email    string
	Username string
	Password string

	User user.User
}

func FromRegisterDTO(dto dto.RegisterData, hashedPass string) Credentials {
	return Credentials{
		Email:    dto.Email,
		Username: dto.Username,
		Password: hashedPass,
	}
}
