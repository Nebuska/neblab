package dto

type RegisterData struct {
	Username string // Max 30 Min 3 unique
	Password string // Max 72byte
	Email    string // email

	FirstName string // 3-30 char
	LastName  string // 3-30 char
}
