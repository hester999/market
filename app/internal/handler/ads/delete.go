package ads

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"market/app/internal/apperr"
	"market/app/internal/handler/ads/dto"
	"net/http"
)

// Delete godoc
// @Summary      Удалить объявление
// @Description  Удаляет объявление по ID. Только владелец может удалить объявление. Требует авторизации.
// @Tags         ads
// @Security     BearerAuth
// @Param        id   path      string  true  "ID объявления"
// @Success      200  {string}  string  "OK"
// @Failure      400  {object}  dto.ErrResponse400 "ID объявления не передан"
// @Failure      401  {object}  dto.ErrResponse401  "Пользователь не авторизован"
// @Failure      403  {object}  dto.ErrResponse403  "Нет прав на удаление чужого объявления"
// @Failure      404  {object}  dto.ErrResponse404  "Объявление не найдено"
// @Failure      500  {object}  dto.ErrResponse500  "Внутренняя ошибка сервера"
// @Router       /api/v1/ads/{id} [delete]
func (a *AdsHandler) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	userId, ok := r.Context().Value("user_id").(string)
	if !ok || userId == "" {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(dto.ErrResponse{
			Message: "unauthorized",
			Code:    http.StatusUnauthorized,
		})
		return
	}

	adId := mux.Vars(r)["id"]
	if adId == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(dto.ErrResponse{
			Message: "ad ID is required",
			Code:    http.StatusBadRequest,
		})
		return
	}

	err := a.ads.Delete(adId, userId)
	if err != nil {
		if errors.Is(err, apperr.ErrAdsNotFound) {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(dto.ErrResponse{
				Message: "ad not found",
				Code:    http.StatusNotFound,
			})
			return
		}
		if errors.Is(err, apperr.ErrForbidden) {
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(dto.ErrResponse{
				Message: "you are not the owner of this ad",
				Code:    http.StatusForbidden,
			})
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(dto.ErrResponse{
			Message: "internal server error",
			Code:    http.StatusInternalServerError})
		return
	}

	w.WriteHeader(http.StatusOK)
}
