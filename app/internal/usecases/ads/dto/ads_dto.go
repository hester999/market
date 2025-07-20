package dto

import (
	entity2 "market/app/internal/entity"
)

type AdResponse struct {
	Id          string
	Title       string
	Description string
	Price       float64
	Author      string
	AuthorID    string
	IsOwner     bool
	Images      []string
}

type AdDetailed struct {
	Ad      entity2.Ad
	Author  string
	Images  []entity2.AdImage
	IsOwner bool
}
