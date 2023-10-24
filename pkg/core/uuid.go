package core

import "github.com/google/uuid"

func IsUuid(value *string) bool {
	if value == nil {
		return false
	}
	_, err := uuid.Parse(*value)
	if err == nil {
		return true
	}
	return false
}

func IsNotUuid(value *string) bool {
	return !IsUuid(value)
}
