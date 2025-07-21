package ads_repo

import (
	"database/sql"
	"errors"
	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"market/app/internal/apperr"
	"market/app/internal/entity"

	"time"
)

type AdDTO struct {
	Id          string    `db:"id"`
	Title       string    `db:"title"`
	Description string    `db:"description"`
	Price       float64   `db:"price"`
	CreatedAt   time.Time `db:"created_at"`
	AuthorId    string    `db:"author_id"`
}

type AdsRepository struct {
	db *sqlx.DB
}

func NewAdsRepository(db *sqlx.DB) *AdsRepository {
	return &AdsRepository{db}
}

func (r *AdsRepository) Create(ad entity.Ad) (entity.Ad, error) {
	query := `
		INSERT INTO ads (id, title, description, price, created_at, author_id)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, title, description, price, created_at, author_id;
	`

	var tmp AdDTO
	err := r.db.Get(&tmp, query,
		ad.Id,
		ad.Title,
		ad.Description,
		ad.Price,
		ad.CreatedAt,
		ad.AuthorId,
	)

	savedAd := entity.Ad{
		Id:          tmp.Id,
		Title:       tmp.Title,
		Description: tmp.Description,
		Price:       tmp.Price,
		CreatedAt:   tmp.CreatedAt,
		AuthorId:    tmp.AuthorId,
	}
	return savedAd, err
}

// GetAll — получение всех объявлений с фильтрацией, сортировкой и пагинацией
func (r *AdsRepository) GetAll(limit, offset int, sortBy, order string, priceMin, priceMax float64) ([]entity.Ad, error) {
	if limit == 0 {
		limit = 10
	}

	allowedSortFields := map[string]bool{"created_at": true, "price": true}
	allowedOrder := map[string]bool{"asc": true, "desc": true}

	if !allowedSortFields[sortBy] {
		sortBy = "created_at"
	}
	if !allowedOrder[order] {
		order = "desc"
	}

	psql := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

	query := psql.
		Select("id", "title", "description", "price", "created_at", "author_id").
		From("ads").
		Limit(uint64(limit)).
		Offset(uint64(offset)).
		OrderBy(sortBy + " " + order)

	if priceMin > 0 && priceMax > 0 {
		query = query.Where(squirrel.And{
			squirrel.GtOrEq{"price": priceMin},
			squirrel.LtOrEq{"price": priceMax},
		})
	} else if priceMin > 0 {
		query = query.Where(squirrel.GtOrEq{"price": priceMin})
	} else if priceMax > 0 {
		query = query.Where(squirrel.LtOrEq{"price": priceMax})
	}

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	var tmp []AdDTO
	err = r.db.Select(&tmp, sqlQuery, args...)
	if err != nil {
		return nil, err
	}
	if len(tmp) == 0 {
		return nil, apperr.ErrAdsNotFound
	}

	ads := make([]entity.Ad, 0, len(tmp))
	for _, v := range tmp {
		ads = append(ads, entity.Ad{
			Id:          v.Id,
			Title:       v.Title,
			Description: v.Description,
			Price:       v.Price,
			CreatedAt:   v.CreatedAt,
			AuthorId:    v.AuthorId,
		})
	}
	return ads, nil
}
func (r *AdsRepository) GetById(adId string) (entity.Ad, error) {
	query := `
		SELECT id, title, description, price, created_at, author_id
		FROM ads
		WHERE id = $1
	`
	var tmp AdDTO
	err := r.db.Get(&tmp, query, adId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entity.Ad{}, apperr.ErrAdsNotFound
		}
		return entity.Ad{}, err
	}

	ad := entity.Ad{
		Id:          tmp.Id,
		Title:       tmp.Title,
		Description: tmp.Description,
		Price:       tmp.Price,
		CreatedAt:   tmp.CreatedAt,
		AuthorId:    tmp.AuthorId,
	}
	return ad, nil
}

// Delete — удаляет объявление, если принадлежит userId
func (r *AdsRepository) Delete(userId, adId string) error {
	query := `DELETE FROM ads WHERE id = $1 AND author_id = $2`

	result, err := r.db.Exec(query, adId, userId)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return apperr.ErrAdsNotFound
	}
	return nil
}

// GetAuthorName — получить имя автора по userId
func (r *AdsRepository) GetAuthorName(userId string) (string, error) {
	query := `SELECT username FROM users WHERE id = $1`
	var username string
	err := r.db.Get(&username, query, userId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", apperr.ErrUserNotFound
		}
		return "", err
	}
	return username, nil
}
