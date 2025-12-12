package dto

import "errors"

var (
	ErrBodyBindingFailOnRegister = errors.New("body binding fail on registering")
	ErrBodyBindingFailOnLogin    = errors.New("body binding fail on login")
)
