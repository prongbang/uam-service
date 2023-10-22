package common

import (
	"encoding/hex"
	"encoding/json"
	"reflect"
	"regexp"
)

func BytesToHex(bytes []byte) string {
	return hex.EncodeToString(bytes)
}

func HexToBytes(text string) ([]byte, error) {
	return hex.DecodeString(text)
}

func AnyToStruct[T any](input any) *T {
	if input != nil {
		if data, err := json.Marshal(input); err == nil {
			t := new(T)
			if err = json.Unmarshal(data, t); err == nil {
				return t
			}
		}
	}
	return nil
}

func StructToMap(input any) map[string]any {
	result := make(map[string]any)
	val := reflect.ValueOf(input)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	if val.Kind() != reflect.Struct {
		return result
	}
	typ := val.Type()
	for i := 0; i < val.NumField(); i++ {
		result[typ.Field(i).Name] = val.Field(i).Interface()
	}
	return result
}

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
