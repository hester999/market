package mapper

import (
	"market/app/internal/entity"
	"market/app/internal/handler/ads/dto"
	usecases "market/app/internal/usecases/ads/dto"
)

func ToAdEntity(dto dto.AdsCreateDTO, authorID string) entity.Ad {
	return entity.Ad{
		Title:       dto.Title,
		Description: dto.Description,
		Price:       dto.Price,
		AuthorId:    authorID,
	}
}

func DtoUsecaseResponseToAdResponse(data usecases.AdResponse) dto.AdResponseDTO {

	return dto.AdResponseDTO{
		Id:          data.Id,
		Title:       data.Title,
		Description: data.Description,
		Price:       data.Price,
		AuthorName:  data.Author,
		AuthorId:    data.AuthorID,
		IsOwner:     data.IsOwner,
		ImagesURl:   data.Images,
	}
}

func DtoUsecaseGetToDtoHandler(data []usecases.AdResponse) dto.AdsResponseDTO {
	res := dto.AdsResponseDTO{}
	res.Ads = make([]dto.AdResponseDTO, 0, len(data))
	for _, item := range data {
		res.Ads = append(res.Ads, DtoUsecaseResponseToAdResponse(item))
	}
	return res
}

func ToAdDetailedResponseDTO(data usecases.AdDetailed) dto.AdDetailedResponseDTO {
	images := make([]string, 0, len(data.Images))
	for _, img := range data.Images {
		images = append(images, img.ImageURL)
	}

	return dto.AdDetailedResponseDTO{
		Id:          data.Ad.Id,
		Title:       data.Ad.Title,
		Description: data.Ad.Description,
		Price:       data.Ad.Price,
		CreatedAt:   data.Ad.CreatedAt,
		AuthorId:    data.Ad.AuthorId,
		AuthorName:  data.Author,
		Images:      images,
		IsOwner:     data.IsOwner,
	}
}
