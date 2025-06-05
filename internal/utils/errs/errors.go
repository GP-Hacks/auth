package errs

import "errors"

var (
	NotFoundError      error = errors.New("not found")
	AlreadyExistsError error = errors.New("already exists")
	SomeError          error = errors.New("some")
	UnauthorizedError  error = errors.New("unauthorized")
)
