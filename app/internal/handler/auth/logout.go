package auth

import (
	"encoding/json"
	"log"
	"market/app/internal/handler/auth/dto"
	"net/http"
)

// Logout godoc
// @Summary      Выход пользователя
// @Description  Удаляет сессию пользователя. Требует авторизации (по токену).
// @Tags         auth
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Success      200  "Успешный выход из системы"
// @Failure      401  {object} dto.ErrResponse401Unauthorized
// @Failure      500  {object}  dto.Err500
// @Router       /api/v1/logout [post]
func (a *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userID, ok := r.Context().Value("user_id").(string)
	if !ok || userID == "" {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(dto.ErrDTO{
			Code:    http.StatusUnauthorized,
			Message: "unauthorized",
		})
		return
	}

	err := a.usecase.Logout(userID)
	if err != nil {
		log.Println("Error logout:", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(dto.ErrDTO{
			Code:    http.StatusInternalServerError,
			Message: "log out failed",
		})
		return
	}

	w.WriteHeader(http.StatusOK)
}
