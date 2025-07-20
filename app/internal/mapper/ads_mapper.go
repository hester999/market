package mapper

import (
	"market/internal/dto"
	"market/internal/entity"
)

func CreateAdDtoToEntity(dto dto.AdsCreateDTO) entity.Ads {
	return entity.Ads{
		Title:       dto.Title,
		Description: dto.Description,
		Price:       dto.Price,
	}
}

// currentUserID, authorName string
func AdResponseToDto(ad entity.Ads) dto.AdResponseDTO {
	return dto.AdResponseDTO{
		Id:          ad.Id,
		Title:       ad.Title,
		Price:       ad.Price,
		AuthorId:    ad.AuthorId,
		Created:     ad.Created,
		Description: ad.Description,
		//IsOwner:     ad.AuthorId == currentUserID,
		//AuthorName:  authorName,
	}
}

//	func AdsResponseToDto(ads []entity.Ads, currentUserID string) dto.AdsResponseDTO {
//		var result dto.AdsResponseDTO
//		for _, val := range ads {
//			tmp := AdResponseToDto(val, currentUserID)
//			result.Ads = append(result.Ads, tmp)
//		}
//		return result
//	}
func AdsResponseToDto(ads []entity.Ads) dto.AdsResponseDTO {
	var result dto.AdsResponseDTO
	for _, val := range ads {
		tmp := AdResponseToDto(val)
		result.Ads = append(result.Ads, tmp)
	}
	return result
}
