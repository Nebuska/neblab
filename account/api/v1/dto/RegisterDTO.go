package dto

import "github.com/Nebuska/neblab/account/internal/auth/dto"

type RegisterDTO struct {
	Username string `json:"username" binding:"required,min=3,max=30"`
	Password string `json:"password" binding:"required,min=8,containsany=0123456789,containsany=abcdefghijklmnopqrstuvwxyz,containsany=ABCDEFGHIJKLMNOPQRSTUVWXYZ"`
	Email    string `json:"email" binding:"required,email"`
}

func (r RegisterDTO) ToServiceRegisterData() dto.RegisterData {
	return dto.RegisterData{
		Username: r.Username,
		Password: r.Password,
		Email:    r.Email,
	}
}
