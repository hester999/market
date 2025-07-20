package dto

type ErrResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

type ErrImagesNotFound struct {
	Message string        `json:"message"`
	Code    int           `json:"code"`
	Data    []ResponseDTO `json:"data"`
}

type Err400BadRequest struct {
	Message string `json:"message" example:"bad request"`
	Code    int    `json:"code" example:"400"`
}

type Err401Unauthorized struct {
	Message string `json:"message" example:"unauthorized"`
	Code    int    `json:"code" example:"401"`
}

type Err404AdNotFound struct {
	Message string `json:"message" example:"ad not found"`
	Code    int    `json:"code" example:"404"`
}

type ErrImagesNotFoundExample struct {
	Message string     `json:"message" example:"image not found"`
	Code    int        `json:"code" example:"404"`
	Data    []struct{} `json:"data"`
}

type Err500Internal struct {
	Message string `json:"message" example:"internal server error"`
	Code    int    `json:"code" example:"500"`
}
