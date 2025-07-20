package repo

import (
	"github.com/jmoiron/sqlx"
	"market/app/internal/entity"
	"time"
)

type Reg struct {
	db *sqlx.DB
}

func NewRegistry(db *sqlx.DB) *Reg {
	return &Reg{db: db}
}

func (r *Reg) Registration(user entity.User) (entity.User, error) {
	query := `
		INSERT INTO users (id, username, email, password_hash, created_at)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, username, email, created_at;
	`

	res := struct {
		Id        string    `db:"id"`
		Username  string    `db:"username"`
		Email     string    `db:"email"`
		CreatedAt time.Time `db:"created_at"`
	}{}

	err := r.db.Get(&res, query, user.Id, user.Username, user.Email, user.PasswordHash, user.CreatedAt)
	if err != nil {
		return entity.User{}, err
	}

	newUser := entity.User{
		Id:        res.Id,
		Username:  res.Username,
		Email:     res.Email,
		CreatedAt: res.CreatedAt,
	}

	return newUser, nil
}

func (r *Reg) EmailExists(email string) (bool, error) {
	var exists bool
	query := `SELECT EXISTS (SELECT 1 FROM users WHERE email = $1 LIMIT 1)`
	err := r.db.Get(&exists, query, email)
	if err != nil {
		return false, err
	}
	return exists, nil
}
