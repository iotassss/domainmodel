package domain

import (
	"regexp"

	"github.com/google/uuid"
)

const (
	uuidPattern = `^[0-9a-f]{8}-[0-9a-f]{4}-4[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$`
)

type UUID struct {
	value string
}

func GenerateUUID() (UUID, error) {
	newUUID := uuid.New()
	return UUID{value: newUUID.String()}, nil
}

func NewUUID(s string) (UUID, error) {
	if !isValidUUID(s) {
		return UUID{}, &ValidationError{Msg: "invalid UUID format"}
	}
	return UUID{value: s}, nil
}

func (u UUID) Value() string {
	return u.value
}

func isValidUUID(uuid string) bool {
	matched := regexp.MustCompile(uuidPattern).MatchString(uuid)
	return matched
}
