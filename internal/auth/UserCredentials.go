package auth

import (
	"github.com/Nebuska/task-tracker/internal/user"

	"gorm.io/gorm"
)

type UserCredentials struct {
	gorm.Model
	Username string `gorm:"unique;size:30;not null" validate:"required,min=3,max=30"`
	Password string `gorm:"not null"`

	User user.User `gorm:"foreignKey:ID;references:ID;constraint:OnDelete:CASCADE"`
}
