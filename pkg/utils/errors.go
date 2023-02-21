package utils

import "errors"

var (
	ErrNotFound         = errors.New("record not found")
	ErrAlreadyExists    = errors.New("user with this username and phone already in blacklist")
	ErrWrongCredentials = errors.New("Invalid credentials")
)
