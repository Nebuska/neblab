package Task

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Name        string
	Description string
	Status      string
	BoardID     int
}
