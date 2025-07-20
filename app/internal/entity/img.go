package entity

import "time"

type AdImage struct {
	Id        string
	AdId      string
	ImageURL  string
	CreatedAt time.Time
}
