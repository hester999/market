package mapper

import (
	"market/app/internal/entity"
	"market/app/internal/handler/reg/dto"
)

func RegRequestDTOToEntity(requestDTO dto.RegUserRequestDTO) entity.User {
	return entity.User{
		Username:     requestDTO.Name,
		Email:        requestDTO.Email,
		PasswordHash: requestDTO.Password,
	}
}

func RegResponseEntityToDTO(user entity.User) dto.RegUserResponseDTO {
	return dto.RegUserResponseDTO{
		Id:        user.Id,
		Name:      user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}
}
