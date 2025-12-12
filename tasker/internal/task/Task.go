package task

import (
	_ "github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Name        string `gorm:"size:50;not null" validate:"required,min=3,max=50"`
	Description string `gorm:"type:text" validate:"omitempty,min=5"`
	Status      string `gorm:"size:20" validate:"oneof=Planning InProgress Completed"`
	BoardID     uint   `gorm:"not null"`
}
