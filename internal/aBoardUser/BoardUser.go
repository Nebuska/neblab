package aBoardUser

import "gorm.io/gorm"

type BoardUser struct {
	gorm.Model

	UserID  uint `gorm:"not null;index"`
	BoardID uint `gorm:"not null;index"`
	Role    string
}
