package utils

import "github.com/google/uuid"

func GenerateRandomUUIDString() string {
	return uuid.New().String()
}

func GenerateUUID() uuid.UUID {
	return uuid.New()
}

func ParseUUID(uuidString string) (uuid.UUID, error) {
	return uuid.Parse(uuidString)
}
