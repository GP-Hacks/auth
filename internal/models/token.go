package models

import "time"

type TokenType int

const (
	Access  TokenType = iota
	Refresh TokenType = iota
)

type Token struct {
	ID        int64
	JTI       string
	SubjectID int64
	Type      TokenType
	Revoked   bool
	IssuedAt  time.Time
	ExpiresAt time.Time
}
