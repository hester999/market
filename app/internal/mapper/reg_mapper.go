package mapper

import (
	"market/app/internal/entity"
	"market/internal/dto"
)

func RegRequestDTOToEntity(requestDTO dto.RegUserRequestDTO) entity.User {
	return entity.User{
		Name:         requestDTO.Name,
		Email:        requestDTO.Email,
		PasswordHash: requestDTO.Password,
	}
}

func RegResponseEntityToDTO(user entity.User) dto.RegUserResponseDTO {
	return dto.RegUserResponseDTO{
		Id:        user.Id,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}
}
