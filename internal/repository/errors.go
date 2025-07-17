package repository

import "errors"

var (
	ErrUserExists = errors.New("User with such login already exists")
	ErrUserNotFound = errors.New("No user with such login was found")
)
