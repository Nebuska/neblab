package dto

import "errors"

var (
	ErrUserNotFound         = errors.New("user not found")
	ErrUserAlreadyExist     = errors.New("email already exists")
	ErrWrongPassword        = errors.New("wrong password")
	ErrExpiredSession       = errors.New("expired session")
	ErrSessionNotFound      = errors.New("session not found")
	ErrSessionMismatch      = errors.New("session mismatch")
	ErrSessionTokenConflict = errors.New("session token conflict")
)
