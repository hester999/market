package img

import "market/app/internal/entity"

type Img interface {
	Create(img entity.AdImage) (entity.AdImage, error)
	GetImages(adId string) ([]entity.AdImage, error)
	GetImageById(id string) (entity.AdImage, error)
	Exists(adId string) (bool, error)
}
