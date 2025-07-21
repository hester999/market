package image

import (
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"io"
	"market/app/internal/apperr"
	"market/app/internal/handler/image/dto"
	"market/app/internal/handler/image/mapper"
	"mime/multipart"
	"net/http"
)

type ImageHandler struct {
	img Img
}

func NewImageHandler(img Img) *ImageHandler {
	return &ImageHandler{img}
}

// AddImage godoc
// @Summary      Загрузить изображение для объявления
// @Description  Загружает изображение в формате JPEG или PNG для указанного объявления. Требует авторизации.
// @Tags         image
// @Accept       multipart/form-data
// @Produce      json
// @Security     BearerAuth
// @Param        id     path      string           true  "ID объявления"
// @Param        image  formData  file             true  "Изображение (jpeg или png)"
// @Success      201    {object}  dto.ResponseDTO               "Успешная загрузка изображения"
// @Failure      400    {object}  dto.Err400BadRequest          "Невалидный файл или данные"
// @Failure      401    {object}  dto.Err401Unauthorized        "Пользователь не авторизован"
// @Failure      404    {object}  dto.Err404AdNotFound          "Объявление не найдено"
// @Failure      415    {object}  dto.Err415                        "Неподдерживаемый тип файла"
// @Failure      500    {object}  dto.Err500Internal            "Внутренняя ошибка сервера"
// @Router       /api/v1/ads/{id}/images [post]
func (i *ImageHandler) AddImage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	adID := mux.Vars(r)["id"]

	if err := uuid.Validate(adID); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(dto.ErrResponse{
			Message: "invalid ad id",
			Code:    http.StatusBadRequest,
		})
		return
	}

	if err := r.ParseMultipartForm(10 << 20); err != nil {

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(dto.ErrResponse{
			Code:    http.StatusBadRequest,
			Message: "invalid request body",
		})
		return
	}

	file, _, err := r.FormFile("image")
	if err != nil {

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(dto.ErrResponse{
			Code:    http.StatusInternalServerError,
			Message: "image upload error",
		})
		return
	}
	defer file.Close()

	ext, err := validateType(file)

	if err != nil {

		if errors.Is(err, apperr.ErrUnsupportedFileType) {
			w.WriteHeader(http.StatusUnsupportedMediaType)
			json.NewEncoder(w).Encode(dto.ErrResponse{
				Message: "unsupported file type",
				Code:    http.StatusUnsupportedMediaType,
			})
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(dto.ErrResponse{
			Message: "bad request",
			Code:    http.StatusBadRequest,
		})
		return
	}

	imgBytes, err := io.ReadAll(file)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(dto.ErrResponse{
			Code:    http.StatusInternalServerError,
			Message: "internal server error",
		})
		return
	}
	res, err := i.img.AddImage(adID, imgBytes, ext)
	if err != nil {
		if errors.Is(err, apperr.ErrAddNotFound) {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(dto.ErrResponse{
				Code:    http.StatusNotFound,
				Message: "ad not found",
			})
			return
		}

		json.NewEncoder(w).Encode(dto.ErrResponse{
			Code:    http.StatusInternalServerError,
			Message: "internal server error",
		})
		return
	}
	response := mapper.EntityImageToDTO(res)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

// GetImages godoc
// @Summary      Получить изображения объявления
// @Description  Возвращает список изображений, прикреплённых к объявлению
// @Tags         image
// @Produce      json
// @Param        id   path      string  true  "ID объявления"
// @Success      200  {array}   dto.ResponseDTO                 "Список изображений"
// @Failure      400  {object}  dto.Err400BadRequest            "Некорректный ID"
// @Failure      404  {object}  dto.ErrImagesNotFoundExample           "Изображения не найдены или объявление не существует"
// @Failure      500  {object}  dto.Err500Internal              "Внутренняя ошибка сервера"
// @Router       /api/v1/ads/{id}/images [get]
func (i *ImageHandler) GetImages(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	adID := mux.Vars(r)["id"]

	if adID == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(dto.ErrResponse{
			Code:    http.StatusBadRequest,
			Message: "ad id is required",
		})
		return
	}

	if err := uuid.Validate(adID); err != nil {

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(dto.ErrResponse{
			Message: "invalid ad id",
			Code:    http.StatusBadRequest,
		})
		return
	}

	res, err := i.img.GetImages(adID)
	if err != nil {
		if errors.Is(err, apperr.ErrAddNotFound) {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(dto.ErrResponse{
				Code:    http.StatusNotFound,
				Message: "ad not found",
			})
			return
		}
		if errors.Is(err, apperr.ErrImgNotFound) {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(dto.ErrImagesNotFound{
				Code:    http.StatusNotFound,
				Message: "image not found",
				Data:    make([]dto.ResponseDTO, 0),
			})
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(dto.ErrResponse{
			Code:    http.StatusInternalServerError,
			Message: "internal server error",
		})
		return
	}
	response := mapper.EntityImagesToDTO(res)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)

}

// GetImageById godoc
// @Summary      Получить изображение по ID
// @Description  Возвращает одно изображение по его ID
// @Tags         image
// @Produce      json
// @Param        id   path      string  true  "ID изображения"
// @Success      200  {object}  dto.ResponseDTO                 "Данные изображения"
// @Failure      400  {object}  dto.Err400BadRequest            "ID изображения не указан"
// @Failure      404  {object}  dto.Err404AdNotFound            "Изображение не найдено"
// @Failure      500  {object}  dto.Err500Internal              "Внутренняя ошибка сервера"
// @Router       /api/v1/ads/images/{id} [get]
func (i *ImageHandler) GetImageById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["id"]

	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(dto.ErrResponse{
			Code:    http.StatusBadRequest,
			Message: "image id required",
		})
		return
	}

	if err := uuid.Validate(id); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(dto.ErrResponse{
			Message: "invalid image id",
			Code:    http.StatusBadRequest,
		})
		return
	}

	res, err := i.img.GetImageById(id)
	if err != nil {
		if errors.Is(err, apperr.ErrImgNotFound) {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(dto.ErrResponse{
				Code:    http.StatusNotFound,
				Message: "image not found",
			})
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(dto.ErrResponse{
			Code:    http.StatusInternalServerError,
			Message: "internal server error",
		})
		return
	}

	response := mapper.EntityImageToDTO(res)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func validateType(file multipart.File) (string, error) {
	head := make([]byte, 512)
	n, err := file.Read(head)
	if err != nil {
		return "", err
	}
	head = head[:n]

	contentType := http.DetectContentType(head)

	if contentType != "image/jpeg" && contentType != "image/png" {
		return "", apperr.ErrUnsupportedFileType
	}

	if _, err := file.Seek(0, 0); err != nil {
		return "", err
	}

	return contentType, nil
}
