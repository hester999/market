package reg

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"market/app/internal/apperr"
	"market/app/internal/entity"
	"market/app/internal/utils"
	"net/mail"
	"strings"
	"time"
	"unicode"
)

type RegistryUsecase struct {
	repo Registry
}

func NewRegistry(repo Registry) *RegistryUsecase {
	return &RegistryUsecase{repo}
}

func (r *RegistryUsecase) Registration(user entity.User) (entity.User, error) {
	if err := r.validateEmail(user.Email); err != nil {
		return entity.User{}, fmt.Errorf("email validation failed: %w", err)
	}
	if err := r.validatePassword(user.PasswordHash); err != nil {
		return entity.User{}, fmt.Errorf("password validation failed: %w", err)
	}

	exists, err := r.repo.EmailExists(user.Email)
	if err != nil {
		return entity.User{}, fmt.Errorf("email existence check failed: %w", err)
	}
	if exists {
		return entity.User{}, fmt.Errorf("email already exists: %w", apperr.ErrEmailAlreadyExists)
	}

	hashed, err := hashPassword(user.PasswordHash)
	if err != nil {
		return entity.User{}, fmt.Errorf("password hashing failed: %w", err)
	}
	user.PasswordHash = hashed

	user.Id, err = utils.GenerateUUID()
	if err != nil {
		return entity.User{}, fmt.Errorf("generate uuid failed: %w", err)
	}
	user.CreatedAt = time.Now().UTC()

	return r.repo.Registration(user)
}

func (r *RegistryUsecase) validateEmail(email string) error {
	email = strings.TrimSpace(email)
	_, err := mail.ParseAddress(email)
	if err != nil {
		return apperr.ErrInvalidEmail
	}
	return nil
}

func (r *RegistryUsecase) validatePassword(password string) error {
	if len(password) < 8 {
		return apperr.ErrInvalidLenPassword
	}

	var hasUpper, hasLower, hasDigit, hasSpecial bool
	for _, ch := range password {
		switch {
		case unicode.IsUpper(ch):
			hasUpper = true
		case unicode.IsLower(ch):
			hasLower = true
		case unicode.IsDigit(ch):
			hasDigit = true
		case unicode.IsPunct(ch) || unicode.IsSymbol(ch):
			hasSpecial = true
		}
	}

	if !hasUpper {
		return apperr.ErrNonUpperCharPass
	}
	if !hasLower {
		return apperr.ErrNonLowerCharPass
	}
	if !hasDigit {
		return apperr.ErrNonDigitPass
	}
	if !hasSpecial {
		return apperr.ErrNonSpecialPass
	}
	return nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}
