package dto

type ErrDTO struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

type ErrDTO400 struct {
	Message string `json:"message" example:"invalid json"`
	Code    int    `json:"code" example:"400"`
}

type ErrDTO409 struct {
	Message string `json:"message" example:"email already exists"`
	Code    int    `json:"code" example:"409"`
}

type ErrDTO500 struct {
	Message string `json:"message" example:"internal server error"`
	Code    int    `json:"code" example:"500"`
}
