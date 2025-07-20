package entity

import "time"

type Ad struct {
	Id          string
	Title       string
	Description string
	Price       float64
	CreatedAt   time.Time
	AuthorId    string
}
