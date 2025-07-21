package img_repo

import (
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
	"market/app/internal/apperr"
	"market/app/internal/entity"
	"time"
)

type ImgRepo struct {
	db *sqlx.DB
}

func NewImgRepo(db *sqlx.DB) *ImgRepo {
	return &ImgRepo{db}
}

type dto struct {
	Id        string    `db:"id"`
	AdId      string    `db:"ad_id"`
	ImageUrl  string    `db:"image_url"`
	CreatedAt time.Time `db:"created_at"`
}

func (i *ImgRepo) Create(img entity.AdImage) (entity.AdImage, error) {
	query := `INSERT INTO ad_images (id,ad_id,image_url,created_at) VALUES ($1,$2,$3,$4) RETURNING id,ad_id,image_url,created_at`

	tmp := struct {
		Id        string    `db:"id"`
		AdId      string    `db:"ad_id"`
		ImageUrl  string    `db:"image_url"`
		CreatedAt time.Time `db:"created_at"`
	}{}

	err := i.db.Get(&tmp, query, img.Id, img.AdId, img.ImageURL, img.CreatedAt)
	if err != nil {
		return entity.AdImage{}, err
	}
	res := entity.AdImage{
		Id:        tmp.Id,
		AdId:      tmp.AdId,
		ImageURL:  tmp.ImageUrl,
		CreatedAt: tmp.CreatedAt,
	}
	return res, nil
}

func (i *ImgRepo) Delete(imgID string) error {
	query := `DELETE FROM ad_images WHERE id=$1`

	_, err := i.db.Exec(query, imgID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return apperr.ErrImgNotFound
		}
		return err
	}
	return nil
}

func (i *ImgRepo) GetImages(adId string) ([]entity.AdImage, error) {
	query := `
		SELECT id, ad_id, image_url, created_at
		FROM ad_images
		WHERE ad_id = $1
		ORDER BY created_at;
	`
	var images []entity.AdImage

	var tmp []dto

	err := i.db.Select(&tmp, query, adId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, apperr.ErrImgNotFound
		}
		return nil, err
	}

	for _, v := range tmp {
		i := entity.AdImage{
			Id:        v.Id,
			AdId:      v.AdId,
			ImageURL:  v.ImageUrl,
			CreatedAt: v.CreatedAt,
		}
		images = append(images, i)
	}
	return images, nil
}

func (i *ImgRepo) Exists(adId string) (bool, error) {
	var exists bool
	query := `SELECT EXISTS (SELECT 1 FROM ads WHERE id = $1 LIMIT 1)`

	err := i.db.Get(&exists, query, adId)
	if err != nil {
		return false, err
	}
	return exists, nil
}

func (i *ImgRepo) GetImageById(id string) (entity.AdImage, error) {
	query := `SELECT id, ad_id, image_url, created_at FROM ad_images WHERE id = $1`

	tmp := struct {
		Id        string    `db:"id"`
		AdId      string    `db:"ad_id"`
		ImageUrl  string    `db:"image_url"`
		CreatedAt time.Time `db:"created_at"`
	}{}

	err := i.db.Get(&tmp, query, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entity.AdImage{}, apperr.ErrImgNotFound
		}
		return entity.AdImage{}, err
	}

	res := entity.AdImage{
		Id:        tmp.Id,
		AdId:      tmp.AdId,
		ImageURL:  tmp.ImageUrl,
		CreatedAt: tmp.CreatedAt,
	}
	return res, nil
}
