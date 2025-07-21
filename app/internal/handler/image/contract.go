package image

import "market/app/internal/entity"

type Img interface {
	AddImage(adId string, data []byte, ext string) (entity.AdImage, error)
	GetImages(adId string) ([]entity.AdImage, error)
	GetImageById(id string) (entity.AdImage, error)
}
