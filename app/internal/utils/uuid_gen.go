package utils

import (
	"errors"
	"github.com/google/uuid"
)

func GenerateUUID() (string, error) {

	uuidV4, err := uuid.NewRandom()

	if err != nil {
		return "", errors.New("UUID generation failed")
	}
	return uuidV4.String(), nil
}
