package common

import (
	"reflect"
	"regexp"
)

func Contains[T any](list []T, target T) bool {
	for _, item := range list {
		if reflect.DeepEqual(item, target) {
			return true
		}
	}
	return false
}

func IsEmail(email string) bool {
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	regex, err := regexp.Compile(pattern)
	if err != nil {
		return false
	}
	return regex.MatchString(email)
}
