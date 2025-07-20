package repo

import (
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
	"market/app/internal/apperr"
	entity2 "market/app/internal/entity"
	"time"
)

type Auth struct {
	db *sqlx.DB
}

func NewAuthRepo(db *sqlx.DB) *Auth {
	return &Auth{db}
}

func (a *Auth) GetUserByEmail(email string) (entity2.User, error) {

	query := "SELECT id,username,email,password_hash FROM users WHERE email=$1;"

	res := struct {
		Id       string `db:"id"`
		Username string `db:"username"`
		Email    string `db:"email"`
		Password string `db:"password_hash"`
	}{}

	err := a.db.Get(&res, query, email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entity2.User{}, apperr.ErrEmailNotFound
		}
		return entity2.User{}, err
	}

	userData := entity2.User{
		Id:           res.Id,
		Username:     res.Username,
		Email:        res.Email,
		PasswordHash: res.Password,
	}
	return userData, nil

}

func (a *Auth) CreateSession(session entity2.Session) (entity2.Session, error) {
	query := `
		INSERT INTO sessions (token, user_id, created_at, expires_at)
		VALUES ($1, $2, $3, $4)
		RETURNING token, user_id, created_at, expires_at;
	`

	res := struct {
		Token     string    `db:"token"`
		UserId    string    `db:"user_id"`
		CreatedAt time.Time `db:"created_at"`
		ExpiresAt time.Time `db:"expires_at"`
	}{}

	err := a.db.Get(&res, query, session.Token, session.UserId, session.CreatedAt, session.ExpiresAt)

	if err != nil {

		return entity2.Session{}, err
	}
	newSession := entity2.Session{
		Token:     res.Token,
		UserId:    res.UserId,
		CreatedAt: res.CreatedAt,
		ExpiresAt: res.ExpiresAt,
	}
	return newSession, nil
}

func (a *Auth) UpdateTokenSession(session entity2.Session) (entity2.Session, error) {
	query := "UPDATE sessions SET token = $1,expires_at=$2,created_at =$3 WHERE user_id=$4 RETURNING token, user_id, created_at, expires_at;"

	res := struct {
		Token     string    `db:"token"`
		UserId    string    `db:"user_id"`
		CreatedAt time.Time `db:"created_at"`
		ExpiresAt time.Time `db:"expires_at"`
	}{}

	err := a.db.Get(&res, query,
		session.Token,
		session.ExpiresAt,
		session.CreatedAt,
		session.UserId,
	)
	if err != nil {
		return entity2.Session{}, err
	}
	newSession := entity2.Session{
		Token:     res.Token,
		UserId:    res.UserId,
		CreatedAt: res.CreatedAt,
		ExpiresAt: res.ExpiresAt,
	}

	return newSession, nil

}

func (a *Auth) CheckUserExists(userId string) (bool, error) {
	var exists bool
	query := `SELECT EXISTS (SELECT 1 FROM sessions WHERE user_id = $1 LIMIT 1)`
	err := a.db.Get(&exists, query, userId)
	if err != nil {
		return false, err
	}
	return exists, nil
}

func (a *Auth) FindSession(token string) (entity2.Session, error) {
	query := `
		SELECT token, user_id, created_at, expires_at
		FROM sessions
		WHERE token = $1 AND expires_at > now();
	`

	var res struct {
		Token     string    `db:"token"`
		UserId    string    `db:"user_id"`
		CreatedAt time.Time `db:"created_at"`
		ExpiresAt time.Time `db:"expires_at"`
	}

	err := a.db.Get(&res, query, token)

	if err != nil {

		if errors.Is(err, sql.ErrNoRows) {
			return entity2.Session{}, apperr.ErrSessionExpired
		}
		return entity2.Session{}, err
	}

	return entity2.Session{
		Token:     res.Token,
		UserId:    res.UserId,
		CreatedAt: res.CreatedAt,
		ExpiresAt: res.ExpiresAt,
	}, nil
}

func (a *Auth) DeleteSession(token string) error {
	query := "DELETE FROM sessions WHERE token = $1"
	_, err := a.db.Exec(query, token)
	if err != nil {
		return err
	}
	return nil
}

func (a *Auth) GetSessionByUserId(userId string) (entity2.Session, error) {
	query := `
		SELECT token, user_id, created_at, expires_at
		FROM sessions
		WHERE user_id = $1
	`

	var res struct {
		Token     string    `db:"token"`
		UserId    string    `db:"user_id"`
		CreatedAt time.Time `db:"created_at"`
		ExpiresAt time.Time `db:"expires_at"`
	}

	err := a.db.Get(&res, query, userId)
	if err != nil {

		if errors.Is(err, sql.ErrNoRows) {
			return entity2.Session{}, apperr.ErrSessionExpired
		}
		return entity2.Session{}, err
	}

	return entity2.Session{
		Token:     res.Token,
		UserId:    res.UserId,
		CreatedAt: res.CreatedAt,
		ExpiresAt: res.ExpiresAt,
	}, nil
}
