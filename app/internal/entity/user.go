package entity

import "time"

type User struct {
	Id           string
	Username     string
	Email        string
	PasswordHash string
	CreatedAt    time.Time
}

type UserItems struct {
	User User
	Ad   Ad
}
