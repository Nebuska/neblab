package User

import (
	"task-tracker/internal/BoardUser"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	Username string `gorm:"unique;size:30;not null" validate:"required,min=3,max=30"`
	Password string `gorm:"not null"`

	FirstName string `gorm:"size:30" validate:"omitempty,min=3,max=30"`
	LastName  string `gorm:"size:30" validate:"omitempty,min=3,max=30"`

	Email string `gorm:"unique;not null" validate:"required,email"`

	BoardUser BoardUser.BoardUser `gorm:"foreignKey:UserID"`
}
