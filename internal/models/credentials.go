package models

type Credentials struct {
	ID             int64
	Email          string
	Password       string
	IsVerification bool
}
