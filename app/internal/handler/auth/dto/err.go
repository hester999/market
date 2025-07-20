package dto

type ErrDTO struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

type Err400 struct {
	Message string `json:"message" example:"invalid JSON"`
	Code    int    `json:"code" example:"400"`
}

type Err401 struct {
	Message string `json:"message" example:"incorrect email or password"`
	Code    int    `json:"code" example:"401"`
}

type Err500 struct {
	Message string `json:"message" example:"internal server error"`
	Code    int    `json:"code" example:"500"`
}

type ErrResponse401Unauthorized struct {
	Message string `json:"message" example:"unauthorized"`
	Code    int    `json:"code" example:"401"`
}
