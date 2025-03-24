package utils

import "github.com/google/uuid"

// GenerateUUID gera um novo UUID v4
func GenerateUUID() string {
	return uuid.New().String()
}
