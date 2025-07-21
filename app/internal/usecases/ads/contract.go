package ads

import "market/app/internal/entity"

type AdsRepo interface {
	Create(ad entity.Ad) (entity.Ad, error)
	GetAll(limit, offset int, sortBy, order string, priceMin, priceMax float64) ([]entity.Ad, error)
	GetById(adId string) (entity.Ad, error)
	Delete(userId, adId string) error
	GetAuthorName(userId string) (string, error)
}
