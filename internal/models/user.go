package models

import "time"

type UserStatus string

const (
	DefaultUser UserStatus = "default"
	AdminUser   UserStatus = "admin"
)

type User struct {
	ID          int64
	Email       string
	FirstName   string
	LastName    string
	Surname     string
	DateOfBirth time.Time
	Status      UserStatus
}
