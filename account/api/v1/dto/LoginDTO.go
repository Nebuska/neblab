package dto

import "github.com/Nebuska/neblab/account/internal/auth/dto"

type LoginDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (receiver LoginDTO) ToServiceLoginData() dto.LoginData {
	return dto.LoginData{
		Username: receiver.Username,
		Password: receiver.Password,
	}
}
