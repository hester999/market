package mapper

import (
	"market/app/internal/entity"
	"market/internal/dto"
)

func UserCreateDtoToEntity(dto dto.CreateUserDTO) entity.User {
	return entity.User{
		Email:        dto.Email,
		PasswordHash: dto.Password,
		Name:         dto.Username,
	}
}

func EntityUserResponseToDTO(user entity.User) dto.UserResponseDTO {
	return dto.UserResponseDTO{
		Id:         user.Id,
		Username:   user.Name,
		Email:      user.Email,
		CreateDate: user.CreatedAt,
	}
}
