package user

import (
	"github.com/Nebuska/task-tracker/internal/boardUser"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	FirstName string `gorm:"size:30" validate:"omitempty,min=3,max=30"`
	LastName  string `gorm:"size:30" validate:"omitempty,min=3,max=30"`

	Email string `gorm:"unique;not null" validate:"required,email"`

	BoardUser boardUser.BoardUser `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
}
