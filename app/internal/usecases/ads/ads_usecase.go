package ads

import (
	"fmt"
	"market/app/internal/apperr"
	"market/app/internal/entity"
	"market/app/internal/usecases/ads/dto"
	"market/app/internal/utils"
	"time"
)

type ImgRepo interface {
	GetImages(adId string) ([]entity.AdImage, error)
}
type Ads struct {
	repo AdsRepo
	img  ImgRepo
}

func NewAds(repo AdsRepo, img ImgRepo) *Ads {
	return &Ads{repo, img}
}

func (a *Ads) Create(ad entity.Ad) (entity.Ad, error) {

	id, err := utils.GenerateUUID()
	if err != nil {
		return entity.Ad{}, fmt.Errorf("uuid generation error: %w", err)
	}

	ad.Id = id
	ad.CreatedAt = time.Now().UTC()

	savedAd, err := a.repo.Create(ad)

	if err != nil {
		return entity.Ad{}, fmt.Errorf("ads creation failed: %w", err)
	}

	return savedAd, nil
}

func (a *Ads) GetById(adId, userId string) (dto.AdDetailed, error) {
	ad, err := a.repo.GetById(adId)
	if err != nil {
		return dto.AdDetailed{}, fmt.Errorf("get by id failed: %w", err)
	}

	authorName, err := a.repo.GetAuthorName(ad.AuthorId)
	if err != nil {
		return dto.AdDetailed{}, fmt.Errorf("get author failed: %w", err)
	}

	images, err := a.img.GetImages(ad.Id)
	if err != nil {
		return dto.AdDetailed{}, fmt.Errorf("get images failed: %w", err)
	}

	return dto.AdDetailed{
		Ad:      ad,
		Author:  authorName,
		Images:  images,
		IsOwner: ad.AuthorId == userId,
	}, nil
}

func (a *Ads) Delete(adId, userId string) error {
	ad, err := a.repo.GetById(adId)
	if err != nil {
		return fmt.Errorf("get ad by id failed: %w", err)
	}

	if ad.AuthorId != userId {
		return apperr.ErrForbidden
	}

	if err := a.repo.Delete(userId, adId); err != nil {
		return fmt.Errorf("delete ad failed: %w", err)
	}

	return nil
}
func (a *Ads) GetAll(userId string, limit, offset int, sortBy, order string, priceMin, priceMax float64) ([]dto.AdResponse, error) {

	ads, err := a.repo.GetAll(limit, offset, sortBy, order, priceMin, priceMax)
	if err != nil {

		return nil, fmt.Errorf("get all failed: %w", err)
	}

	var result []dto.AdResponse
	for _, ad := range ads {
		authorName, err := a.repo.GetAuthorName(ad.AuthorId)
		if err != nil {

			return nil, fmt.Errorf("author fetch failed: %w", err)
		}

		images, err := a.img.GetImages(ad.Id)
		if err != nil {

			return nil, fmt.Errorf("get images failed: %w", err)
		}

		imageURLs := make([]string, 0, len(images))
		for _, img := range images {
			imageURLs = append(imageURLs, img.ImageURL)
		}

		dto := dto.AdResponse{
			Id:          ad.Id,
			Title:       ad.Title,
			Description: ad.Description,
			Price:       ad.Price,
			Author:      authorName,
			AuthorID:    ad.AuthorId,
			IsOwner:     ad.AuthorId == userId,
			Images:      imageURLs,
		}

		result = append(result, dto)
	}

	return result, nil
}

func (a *Ads) validateGetAll(limit, offset int, priceMin, priceMax float64) error {
	if limit <= 0 {
		return apperr.ErrInvalidLimit
	}
	if offset < 0 {
		return apperr.ErrInvalidOffset
	}
	if priceMin < 0 || priceMax < 0 || priceMax < priceMin {
		return apperr.ErrInvalidPrice
	}
	return nil
}
