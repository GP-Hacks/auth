package models

import "time"

type User struct {
	ID          int64
	Email       string
	FirstName   string
	LastName    string
	Surname     string
	DateOfBirth time.Time
}
