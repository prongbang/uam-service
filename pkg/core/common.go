package core

import "strings"

func Commas(args []any) string {
	placeholders := make([]string, len(args))
	for i := range args {
		placeholders[i] = "?"
	}
	return strings.Join(placeholders, ", ")
}
