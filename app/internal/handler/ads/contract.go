package ads

import (
	"market/app/internal/entity"
	"market/app/internal/usecases/ads/dto"
)

type Ads interface {
	Create(ad entity.Ad) (entity.Ad, error)
	GetById(adId, userId string) (dto.AdDetailed, error)
	Delete(adId, userId string) error
	GetAll(userId string, limit, offset int, sortBy, order string, priceMin, priceMax float64) ([]dto.AdResponse, error)
}
