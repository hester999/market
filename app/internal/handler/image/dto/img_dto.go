package dto

import "time"

type ResponseDTO struct {
	Id        string    `json:"id" example:"f8c7e2a1-72e1-4c9e-bd84-7ae1b8fc4d4b"`
	AdId      string    `json:"adId" example:"92d1b029-10b6-4df4-8463-b3272e4f15ee"`
	ImageURL  string    `json:"imageUrl" example:"/static/upload/example.jpg"`
	CreatedAt time.Time `json:"createdAt" example:"2025-07-20T12:34:56Z"`
}

type ImagesResponseDTO struct {
	Images []ResponseDTO `json:"images" swaggertype:"array,object"`
}
