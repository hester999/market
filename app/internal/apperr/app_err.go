package apperr

import "errors"

var ErrInvalidUUID = errors.New("uuid validation failed")

// ads err repo
var ErrAdsNotFound = errors.New("ads not found")

// ads err usecases
var (
	ErrTitleTooLong       = errors.New("title is too long")
	ErrDescriptionTooLong = errors.New("description is too long")
	ErrInvalidPrice       = errors.New("price is invalid")
	ErrInvalidLimit       = errors.New("limit is invalid")
	ErrInvalidOffset      = errors.New("offset is invalid")
	ErrForbidden          = errors.New("user is not owner")
)

//reg err

var (
	ErrInvalidLenPassword = errors.New("password must be at least 8 characters long")
	ErrNonUpperCharPass   = errors.New("password must contain at least one uppercase letter")
	ErrNonDigitPass       = errors.New("password must contain at least one digit")
	ErrNonLowerCharPass   = errors.New("password must contain at least one lowercase letter")
	ErrNonSpecialPass     = errors.New("password must contain at least one special character")
	ErrInvalidEmail       = errors.New("email is invalid")
	ErrEmailAlreadyExists = errors.New("email already exists")
	ErrUserNotFound       = errors.New("user not found")
)

// ImgErr
var (
	ErrImgNotFound         = errors.New("image not found")
	ErrAddNotFound         = errors.New("add  not found")
	ErrUnsupportedFileType = errors.New("unsupported file type")
)

// auth err
var (
	ErrEmailNotFound     = errors.New("email not found")
	ErrSessionExpired    = errors.New("session expired")
	ErrIncorrectPassword = errors.New("incorrect password")
)

// authHandlerValidatorErr
var (
	ErrPassRequired  = errors.New("password is required")
	ErrEmailRequired = errors.New("email is required")
)

// regHandlerErr
var (
	ErrNameIsRequired     = errors.New("name is required")
	ErrPasswordIsRequired = errors.New("password is required")
	ErrEmailIsRequired    = errors.New("email is required")
)
