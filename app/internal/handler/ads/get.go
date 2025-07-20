package ads

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"market/app/internal/apperr"
	dto2 "market/app/internal/handler/ads/dto"
	"market/app/internal/handler/ads/mapper"
	"net/http"
	"strconv"
)

// GetAllAds godoc
// @Summary      Получить все объявления
// @Description  Возвращает список всех объявлений. Не требует авторизации, но если токен передан — отмечает ваши объявления как `is_owner=true`.
// @Tags         ads
// @Accept       json
// @Produce      json
// @Param        limit    query     int     false  "Ограничение по количеству"
// @Param        offset   query     int     false  "Смещение"
// @Param        sort     query     string  false  "Поле для сортировки"
// @Param        order    query     string  false  "asc или desc"
// @Param        min      query     number  false  "Минимальная цена"
// @Param        max      query     number  false  "Максимальная цена"
// @Success      200  {object}  dto.AdsResponseDTO
// @Failure      400  {object}  dto.ErrResponse400
// @Failure      404  {object}  dto.ErrResponse404ArrExample
// @Failure      500  {object}  dto.ErrResponse500
// @Router       /api/v1/ads [get]
func (a *AdsHandler) GetAllAds(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var userID string

	if id, ok := r.Context().Value("user_id").(string); ok {
		userID = id
	}

	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	offset, _ := strconv.Atoi(r.URL.Query().Get("offset"))
	sort := r.URL.Query().Get("sort")
	order := r.URL.Query().Get("order")
	minPrice, _ := strconv.ParseFloat(r.URL.Query().Get("min"), 64)
	maxPrice, _ := strconv.ParseFloat(r.URL.Query().Get("max"), 64)

	if err := a.validateQueryParams(limit, offset, minPrice, maxPrice); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(dto2.ErrResponse{
			Message: err.Error(),
			Code:    http.StatusBadRequest,
		})
		return
	}

	res, err := a.ads.GetAll(userID, limit, offset, sort, order, minPrice, maxPrice)
	if err != nil {
		if errors.Is(err, apperr.ErrAdsNotFound) {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(dto2.ErrResponseNotFound{
				Code:    http.StatusNotFound,
				Message: "ads not found",
				Data:    make([]dto2.AdResponseDTO, 0),
			})
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(dto2.ErrResponse{
			Code:    http.StatusInternalServerError,
			Message: "internal server error",
		})
		return
	}
	response := mapper.DtoUsecaseGetToDtoHandler(res)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (a *AdsHandler) validateQueryParams(limit, offset int, priceMin, priceMax float64) error {
	// если явно задан limit, он должен быть > 0
	if limit < 0 {
		return apperr.ErrInvalidLimit
	}
	// offset может быть 0, но не меньше
	if offset < 0 {
		return apperr.ErrInvalidOffset
	}

	// если оба заданы — проверяем корректность диапазона
	if priceMin > 0 && priceMax > 0 && priceMax < priceMin {
		return apperr.ErrInvalidPrice
	}

	// если только один задан — валидировать только его
	if priceMin < 0 || priceMax < 0 {
		return apperr.ErrInvalidPrice
	}

	return nil
}

// GetAdByID godoc
// @Summary      Получить объявление по ID
// @Description  Возвращает детальное объявление. Можно передать токен, чтобы узнать `is_owner`.
// @Tags         ads
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "ID объявления"
// @Success      200  {object}  dto.AdDetailedResponseDTO
// @Failure      400  {object}  dto.ErrResponse400
// @Failure 404 {object} dto.ErrResponse404
// @Failure      500  {object}  dto.ErrResponse500
// @Router       /api/v1/ads/{id} [get]
func (a *AdsHandler) GetAdByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	adId := mux.Vars(r)["id"]
	var userId string
	if id, ok := r.Context().Value("user_id").(string); ok {
		userId = id
	}

	if adId == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(dto2.ErrResponse{
			Code:    http.StatusBadRequest,
			Message: "id is empty",
		})
		return
	}

	res, err := a.ads.GetById(adId, userId)
	if err != nil {
		if errors.Is(err, apperr.ErrAdsNotFound) {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(dto2.ErrResponseNotFound{
				Code:    http.StatusNotFound,
				Message: "ads not found",
			})
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(dto2.ErrResponse{
			Code:    http.StatusInternalServerError,
			Message: "internal server error",
		})
		return

	}

	response := mapper.ToAdDetailedResponseDTO(res)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)

}
