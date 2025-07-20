package reg

import (
	"encoding/json"
	"errors"
	"market/app/internal/apperr"
	"market/app/internal/entity"
	dto2 "market/app/internal/handler/reg/dto"
	"market/app/internal/handler/reg/mapper"
	"net/http"
)

type Registry interface {
	Registration(user entity.User) (entity.User, error)
}

type RegistryHandler struct {
	reg Registry
}

func NewRegistryHandler(reg Registry) *RegistryHandler {
	return &RegistryHandler{reg}
}

// RegistrationHandler godoc
// @Summary      Регистрация пользователя
// @Description  Регистрирует нового пользователя по имени, email и паролю
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        user  body      dto.RegUserRequestDTO     true  "Данные для регистрации"
// @Success      201   {object}  dto.RegUserResponseDTO     "Пользователь успешно зарегистрирован"
// @Failure      400   {object}  dto.ErrDTO400              "Некорректный JSON или поля"
// @Failure      409   {object}  dto.ErrDTO409              "Email уже существует"
// @Failure      500   {object}  dto.ErrDTO500              "Внутренняя ошибка сервера"
// @Router       /api/v1/register [post]
func (reg *RegistryHandler) RegistrationHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var newUser dto2.RegUserRequestDTO

	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(dto2.ErrDTO{Message: "invalid json",
			Code: http.StatusBadRequest})
		return
	}
	err = reg.validateRequest(newUser)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(dto2.ErrDTO{Message: err.Error(), Code: http.StatusBadRequest})
		return
	}
	entityUser := mapper.RegRequestDTOToEntity(newUser)

	result, err := reg.reg.Registration(entityUser)

	if err != nil {
		msgErr, ok := reg.badRequestErr(err)
		if ok {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(dto2.ErrDTO{Message: msgErr.Error(), Code: http.StatusBadRequest})
			return
		}
		if errors.Is(err, apperr.ErrEmailAlreadyExists) {
			w.WriteHeader(http.StatusConflict)
			json.NewEncoder(w).Encode(dto2.ErrDTO{Message: err.Error(), Code: http.StatusConflict})
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(dto2.ErrDTO{Message: "internal server error", Code: http.StatusInternalServerError})
		return
	}

	response := mapper.RegResponseEntityToDTO(result)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)

}

func (reg *RegistryHandler) validateRequest(user dto2.RegUserRequestDTO) error {

	if user.Email == "" && user.Password == "" && user.Name == "" {
		return errors.New("email,name and password are required")
	}

	if user.Name == "" {
		return apperr.ErrNameIsRequired
	}
	if user.Password == "" {
		return apperr.ErrPasswordIsRequired
	}
	if user.Email == "" {
		return apperr.ErrEmailIsRequired
	}
	return nil
}

func (reg *RegistryHandler) badRequestErr(err error) (error, bool) {
	switch {
	case errors.Is(err, apperr.ErrInvalidLenPassword):
		return apperr.ErrInvalidLenPassword, true
	case errors.Is(err, apperr.ErrNonUpperCharPass):
		return apperr.ErrNonUpperCharPass, true
	case errors.Is(err, apperr.ErrNonLowerCharPass):
		return apperr.ErrNonLowerCharPass, true
	case errors.Is(err, apperr.ErrNonDigitPass):
		return apperr.ErrNonDigitPass, true
	case errors.Is(err, apperr.ErrNonSpecialPass):
		return apperr.ErrNonSpecialPass, true
	case errors.Is(err, apperr.ErrInvalidEmail):
		return apperr.ErrInvalidEmail, true
	}

	return nil, false
}
