package mapper

import (
	"market/app/internal/entity"
	"market/app/internal/handler/image/dto"
)

func EntityImageToDTO(image entity.AdImage) dto.ResponseDTO {
	return dto.ResponseDTO{
		Id:        image.Id,
		AdId:      image.AdId,
		ImageURL:  image.ImageURL,
		CreatedAt: image.CreatedAt,
	}
}

func EntityImagesToDTO(images []entity.AdImage) dto.ImagesResponseDTO {

	var res dto.ImagesResponseDTO

	res.Images = make([]dto.ResponseDTO, 0, len(images))

	for _, image := range images {
		res.Images = append(res.Images, EntityImageToDTO(image))
	}
	return res
}
