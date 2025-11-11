package aBoard

import (
	"github.com/Nebuska/task-tracker/internal/aBoardUser"
	"github.com/Nebuska/task-tracker/internal/aTask"

	_ "github.com/go-playground/validator/v10"

	"gorm.io/gorm"
)

type Board struct {
	gorm.Model
	Name       string                 `gorm:"size:30;not null" validate:"required,min=3,max=30"`
	Tasks      []aTask.Task           `gorm:"foreignKey:BoardID;constraint:OnDelete:CASCADE"`
	BoardUsers []aBoardUser.BoardUser `gorm:"foreignKey:BoardID;constraint:OnDelete:CASCADE"`
}
