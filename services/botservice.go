package services

import (
	"github.com/google/uuid"
)

func GenerateUUID() (string, error) {
	// Generate a new UUID (Version 4)
	id, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	return id.String(), nil
}
