package dto

type ErrResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

type ErrResponseNotFound struct {
	Message string          `json:"message"`
	Code    int             `json:"code"`
	Data    []AdResponseDTO `json:"data"`
}

type ErrResponse400 struct {
	Message string `json:"message" example:"Bad Request"`
	Code    int    `json:"code" example:"400"`
}

type ErrResponse401 struct {
	Message string `json:"message" example:"unauthorized"`
	Code    int    `json:"code" example:"401"`
}

type ErrResponse500 struct {
	Message string `json:"message" example:"internal server error"`
	Code    int    `json:"code" example:"500"`
}

// swagger:response ErrResponse404ArrExample
type ErrResponse404ArrExample struct {
	// in: body
	Body struct {
		Message string     `json:"message" example:"Not Found"`
		Code    int        `json:"code" example:"404"`
		Data    []struct{} `json:"data"`
	}
}

type ErrResponse404 struct {
	Message string `json:"message" example:"Not Found"`
	Code    int    `json:"code" example:"404"`
}

type ErrResponse403 struct {
	Message string `json:"message" example:"you are not owner"`
	Code    int    `json:"code" example:"403"`
}
