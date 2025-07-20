package dto

import "time"

type AdsCreateDTO struct {
	Title       string  `json:"title" example:"Велосипед"`
	Description string  `json:"description" example:"Горный велосипед в хорошем состоянии"`
	Price       float64 `json:"price" example:"5000"`
}

type AdResponseDTO struct {
	Id          string    `json:"id" example:"b72f25de-3ef1-4a36-9675-df31545fa08c"`
	Title       string    `json:"title" example:"Велосипед"`
	Description string    `json:"description" example:"Горный велосипед в хорошем состоянии"`
	Price       float64   `json:"price" example:"5000"`
	Created     time.Time `json:"created_at" example:"2025-07-20T12:34:56Z"`
	AuthorId    string    `json:"author_id" example:"a491c857-dbd0-4a4a-88dc-123456789abc"`
	AuthorName  string    `json:"author_name" example:"Иван"`
	IsOwner     bool      `json:"is_owner" example:"true"`
	ImagesURl   []string  `json:"images" example:"['/static/upload/1.jpg','/static/upload/2.png']"`
}

type AdsResponseDTO struct {
	Ads []AdResponseDTO `json:"ads"`
}

type AdCreateRespDTO struct {
	Id          string    `json:"id" example:"b72f25de-3ef1-4a36-9675-df31545fa08c"`
	Title       string    `json:"title" example:"Велосипед"`
	Description string    `json:"description" example:"Горный велосипед в хорошем состоянии"`
	Price       float64   `json:"price" example:"5000"`
	CreatedAt   time.Time `json:"created_at" example:"2025-07-20T12:34:56Z"`
	AuthorId    string    `json:"author_id" example:"a491c857-dbd0-4a4a-88dc-123456789abc"`
}

type AdDetailedResponseDTO struct {
	Id          string    `json:"id" example:"b72f25de-3ef1-4a36-9675-df31545fa08c"`
	Title       string    `json:"title" example:"Велосипед"`
	Description string    `json:"description" example:"Горный велосипед в хорошем состоянии"`
	Price       float64   `json:"price" example:"5000"`
	CreatedAt   time.Time `json:"created_at" example:"2025-07-20T12:34:56Z"`
	AuthorId    string    `json:"author_id" example:"a491c857-dbd0-4a4a-88dc-123456789abc"`
	AuthorName  string    `json:"author_name" example:"Иван"`
	Images      []string  `json:"images" example:"['/static/upload/1.jpg','/static/upload/2.png']"`
	IsOwner     bool      `json:"is_owner" example:"true"`
}
