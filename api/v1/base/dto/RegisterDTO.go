package dto

type RegisterDTO struct {
	Username string `json:"username" binding:"required,min=3,max=30"`
	Password string `json:"password" binding:"required,min=8,containsany=0123456789,containsany=abcdefghijklmnopqrstuvwxyz,containsany=ABCDEFGHIJKLMNOPQRSTUVWXYZ"`
	Email    string `json:"email" binding:"required,email"`
}
