package dto

import "time"

type RegUserRequestDTO struct {
	Name     string `json:"name" example:"Иван"`
	Email    string `json:"email" example:"ivan@example.com"`
	Password string `json:"password" example:"secret123"`
}

type RegUserResponseDTO struct {
	Id        string    `json:"id" example:"b6c859e5-5586-4e52-b02a-82678f30a3fa"`
	Name      string    `json:"name" example:"Иван"`
	Email     string    `json:"email" example:"ivan@example.com"`
	CreatedAt time.Time `json:"createdAt" example:"2025-07-20T12:34:56Z"`
}
