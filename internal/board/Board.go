package board

import (
	"github.com/Nebuska/task-tracker/internal/boardUser"
	"github.com/Nebuska/task-tracker/internal/task"

	_ "github.com/go-playground/validator/v10"

	"gorm.io/gorm"
)

type Board struct {
	gorm.Model
	Name       string                `gorm:"size:30;not null" validate:"required,min=3,max=30"`
	Tasks      []task.Task           `gorm:"foreignKey:BoardID;constraint:OnDelete:CASCADE"`
	BoardUsers []boardUser.BoardUser `gorm:"foreignKey:BoardID;constraint:OnDelete:CASCADE"`
}
