package entity

import "time"

type Session struct {
	Token     string
	UserId    string
	CreatedAt time.Time
	ExpiresAt time.Time
}
