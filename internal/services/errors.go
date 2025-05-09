package services

import "errors"

var (
	NotFound           = errors.New("Not found")
	InternalServer     = errors.New("Internal server")
	AlreadyExists      = errors.New("Already exists")
	InvalidCredentials = errors.New("Invalid credentials")
	InvalidToken       = errors.New("Invalid token")
)
