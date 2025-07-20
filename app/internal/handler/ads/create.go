package ads

import (
	"encoding/json"
	"market/app/internal/apperr"
	dto2 "market/app/internal/handler/ads/dto"
	"market/app/internal/handler/ads/mapper"
	"net/http"
	"unicode/utf8"
)

// Create godoc
// @Summary      Создать объявление
// @Description  Создает новое объявление. Требует авторизации.
// @Tags         ads
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        ad  body  dto.AdsCreateDTO  true  "Создаваемое объявление"
// @Success      201  {object}  dto.AdCreateRespDTO
// @Failure      400  {object}  dto.ErrResponse400
// @Failure      401  {object}  dto.ErrResponse401
// @Failure      500  {object}  dto.ErrResponse500
// @Router       /api/v1/ads [post]
func (a *AdsHandler) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	userID, ok := r.Context().Value("user_id").(string)
	if !ok || userID == "" {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(dto2.ErrResponse{
			Code:    http.StatusUnauthorized,
			Message: "unauthorized",
		})
		return
	}

	var newAd dto2.AdsCreateDTO
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&newAd); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(dto2.ErrResponse{
			Code:    http.StatusBadRequest,
			Message: "invalid JSON",
		})

		return
	}

	err := a.createValidate(newAd)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(dto2.ErrResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	adEntity := mapper.ToAdEntity(newAd, userID)

	createdAd, err := a.ads.Create(adEntity)
	if err != nil {
		json.NewEncoder(w).Encode(dto2.ErrResponse{
			Code:    http.StatusInternalServerError,
			Message: "internal server error",
		})
		return
	}

	resp := dto2.AdCreateRespDTO{
		Id:          createdAd.Id,
		Title:       createdAd.Title,
		Description: createdAd.Description,
		Price:       createdAd.Price,
		CreatedAt:   createdAd.CreatedAt,
		AuthorId:    createdAd.AuthorId,
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}

func (a *AdsHandler) createValidate(ads dto2.AdsCreateDTO) error {
	if utf8.RuneCountInString(ads.Title) > 50 {
		return apperr.ErrTitleTooLong
	}
	if utf8.RuneCountInString(ads.Description) > 1000 {
		return apperr.ErrDescriptionTooLong
	}

	if ads.Price <= 0 {
		return apperr.ErrInvalidPrice
	}
	return nil
}
