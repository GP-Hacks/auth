package services

import "errors"

var (
	NotFound       = errors.New("Not found")
	InternalServer = errors.New("Internal server")
	AlreadyExists  = errors.New("Already exists")
)
