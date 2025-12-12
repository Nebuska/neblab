package credentials

import (
	"github.com/Nebuska/neblab/account/internal/user"

	"gorm.io/gorm"
)

type Credentials struct {
	gorm.Model
	Email    string `gorm:"unique"`
	Username string `gorm:"size:30;not null" validate:"required,min=3,max=30"`
	Password string `gorm:"not null"`

	User user.User `gorm:"foreignKey:ID;references:ID;constraint:OnDelete:CASCADE"`
}
