package auth

import (
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"market/app/internal/apperr"
	"market/app/internal/entity"
	"market/app/internal/utils"
	"net/mail"
	"strings"
	"time"
)

type AuthUsecase struct {
	repo Auth
}

func NewAuth(repo Auth) *AuthUsecase {
	return &AuthUsecase{repo}
}

func (a *AuthUsecase) Login(email, password string) (string, error) {
	if err := a.validateEmail(email); err != nil {
		return "", err
	}

	user, err := a.repo.GetUserByEmail(email)
	if err != nil {
		if errors.Is(err, apperr.ErrEmailNotFound) {
			return "", apperr.ErrUserNotFound
		}
		return "", err
	}

	if err := a.validatePassword(password, user.PasswordHash); err != nil {
		return "", apperr.ErrIncorrectPassword
	}

	token, err := utils.GenerateUUID()
	if err != nil {
		return "", fmt.Errorf("token generation failed: %w", err)
	}

	exists, err := a.repo.CheckUserExists(user.Id)
	if err != nil {
		return "", err
	}

	if !exists {
		session := entity.Session{
			Token:     token,
			UserId:    user.Id,
			CreatedAt: time.Now().UTC(),
			ExpiresAt: time.Now().UTC().Add(24 * time.Hour),
		}
		_, err := a.repo.CreateSession(session)
		if err != nil {
			return "", err
		}
		return token, nil
	}

	// Сессия есть — возвращаем текущий токен (без обновления)
	session, err := a.repo.GetSessionByUserId(user.Id)
	if err != nil {
		return "", err
	}
	return session.Token, nil
}

func (a *AuthUsecase) Logout(userID string) error {

	session, err := a.repo.GetSessionByUserId(userID)

	if err != nil {
		return err
	}
	return a.repo.DeleteSession(session.Token)
}

func (a *AuthUsecase) ValidateSession(token string) (string, error) {
	const bearerPrefix = "Bearer "
	if len(token) > len(bearerPrefix) && strings.HasPrefix(token, bearerPrefix) {
		token = strings.TrimPrefix(token, bearerPrefix)
	}

	session, err := a.repo.FindSession(token)
	if err != nil {

		return "", err
	}

	if session.ExpiresAt.Before(time.Now().UTC()) {
		newToken, err := utils.GenerateUUID()
		if err != nil {

			return "", err
		}

		session.Token = newToken
		session.CreatedAt = time.Now().UTC()
		session.ExpiresAt = session.CreatedAt.Add(24 * time.Hour)

		_, err = a.repo.UpdateTokenSession(session)
		if err != nil {

			return "", err
		}
		return session.UserId, nil
	}

	return session.UserId, nil
}

func (a *AuthUsecase) validateEmail(email string) error {
	email = strings.TrimSpace(email)
	_, err := mail.ParseAddress(email)
	if err != nil {
		return apperr.ErrInvalidEmail
	}
	return nil
}

func (a *AuthUsecase) validatePassword(plainPassword, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
}
