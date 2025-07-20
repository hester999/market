package auth

import (
	"encoding/json"
	"errors"
	"market/app/internal/apperr"
	dto2 "market/app/internal/handler/auth/dto"
	"net/http"
)

// LoginHandler godoc
// @Summary      Вход пользователя
// @Description  Авторизует пользователя по email и паролю и возвращает токен
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        credentials  body  dto.LoginRequestDTO  true  "Данные для входа"
// @Success      200  {object}  dto.LoginResponseDTO
// @Failure      400  {object}  dto.Err400
// @Failure      401  {object}  dto.Err401
// @Failure      500  {object}  dto.Err500
// @Router       /api/v1/login [post]
func (a *AuthHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var reg dto2.LoginRequestDTO
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&reg); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(dto2.ErrDTO{
			Message: "invalid JSON",
			Code:    http.StatusBadRequest,
		})
		return
	}

	err := a.validateStruct(reg)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(dto2.ErrDTO{
			Message: err.Error(),
			Code:    http.StatusBadRequest,
		})
		return
	}

	token, err := a.usecase.Login(reg.Email, reg.Password)
	if err != nil {
		errResp := a.compareErr(err)
		w.WriteHeader(errResp.Code)
		json.NewEncoder(w).Encode(errResp)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(dto2.LoginResponseDTO{Token: token})
	return
}

func (a *AuthHandler) validateStruct(reg dto2.LoginRequestDTO) error {

	if reg.Email == "" && reg.Password == "" {
		return errors.New("email and password are required")
	}
	if reg.Email == "" {

		return apperr.ErrEmailRequired
	}
	if reg.Password == "" {
		return apperr.ErrPassRequired
	}
	return nil
}

func (a *AuthHandler) compareErr(err error) dto2.ErrDTO {
	var res dto2.ErrDTO

	switch {
	case errors.Is(err, apperr.ErrUserNotFound):
		res = dto2.ErrDTO{
			Message: "user not found",
			Code:    http.StatusNotFound,
		}
	case errors.Is(err, apperr.ErrIncorrectPassword):
		res = dto2.ErrDTO{
			Message: "incorrect password",
			Code:    http.StatusUnauthorized,
		}
	case errors.Is(err, apperr.ErrInvalidEmail):
		res = dto2.ErrDTO{
			Message: "invalid email",
			Code:    http.StatusBadRequest,
		}
	default:
		res = dto2.ErrDTO{
			Message: "internal error",
			Code:    http.StatusInternalServerError,
		}
	}

	return res
}
