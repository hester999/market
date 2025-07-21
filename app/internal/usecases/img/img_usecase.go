package img

import (
	"market/app/internal/apperr"
	"market/app/internal/entity"
	"market/app/internal/utils"
	"os"
	"path/filepath"

	"time"
)

type ImgUsecase struct {
	repo Img
}

func NewImgUsecase(repo Img) *ImgUsecase {
	return &ImgUsecase{repo}
}

func (i *ImgUsecase) AddImage(adId string, data []byte, ext string) (entity.AdImage, error) {
	exists, err := i.repo.Exists(adId)
	if err != nil {
		return entity.AdImage{}, err
	}
	if !exists {
		return entity.AdImage{}, apperr.ErrAddNotFound
	}

	imgId, err := utils.GenerateUUID()
	if err != nil {
		return entity.AdImage{}, err
	}
	filename, err := utils.GenerateUUID()
	if err != nil {
		return entity.AdImage{}, err

	}
	ext = i.getExt(ext)
	filename += ext
	imageURL, err := i.saveFile(filename, data)
	if err != nil {

		return entity.AdImage{}, err
	}

	res := entity.AdImage{
		Id:        imgId,
		AdId:      adId,
		ImageURL:  imageURL,
		CreatedAt: time.Now().UTC(),
	}

	res, err = i.repo.Create(res)
	if err != nil {
		return entity.AdImage{}, err
	}
	return res, nil
}

func (i *ImgUsecase) saveFile(filename string, data []byte) (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	uploadPath := filepath.Join(wd, "..", "static", "upload")
	//uploadPath := filepath.Join(wd, "static", "upload")

	if err := os.MkdirAll(uploadPath, 0755); err != nil {
		return "", err
	}

	filepath := filepath.Join(uploadPath, filename)
	file, err := os.Create(filepath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	if _, err := file.Write(data); err != nil {
		return "", err
	}

	publicPath := "/static/upload/" + filename
	return publicPath, nil
}

func (i *ImgUsecase) GetImages(adId string) ([]entity.AdImage, error) {
	exists, err := i.repo.Exists(adId)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, apperr.ErrAddNotFound
	}

	res, err := i.repo.GetImages(adId)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (i *ImgUsecase) GetImageById(
	id string) (entity.AdImage, error) {

	res, err := i.repo.GetImageById(id)
	if err != nil {
		return entity.AdImage{}, err
	}
	return res, nil
}

func (i *ImgUsecase) getExt(ext string) string {

	switch ext {
	case "image/jpeg":
		return ".jpg"
	case "image/png":
		return ".png"
	default:
		return ".bin"
	}
}
