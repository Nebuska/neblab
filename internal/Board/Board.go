package Board

import (
	"task-tracker/internal/Task"

	"gorm.io/gorm"
)

type Board struct {
	gorm.Model
	Name  string
	Tasks []Task.Task `gorm:"foreignKey:BoardID"`
}
