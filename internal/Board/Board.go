package Board

import (
	"task-tracker/internal/BoardUser"
	"task-tracker/internal/Task"

	_ "github.com/go-playground/validator/v10"

	"gorm.io/gorm"
)

type Board struct {
	gorm.Model
	Name       string                `gorm:"size:30;not null" validate:"required,min=3,max=30"`
	Tasks      []Task.Task           `gorm:"foreignKey:BoardID;constraint:OnDelete:CASCADE"`
	BoardUsers []BoardUser.BoardUser `gorm:"foreignKey:BoardID;constraint:OnDelete:CASCADE"`
}
