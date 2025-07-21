package auth

import "market/app/internal/entity"

type Auth interface {
	GetUserByEmail(email string) (entity.User, error)
	CreateSession(session entity.Session) (entity.Session, error)
	UpdateTokenSession(session entity.Session) (entity.Session, error)
	CheckUserExists(userId string) (bool, error)
	FindSession(token string) (entity.Session, error)
	DeleteSession(token string) error
	GetSessionByUserId(id string) (entity.Session, error)
}
