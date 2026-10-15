package tolower

import (
	"unicode"
)

func ToLower(s string) string {
	serierunes := []rune(s)
	for i, r := range serierunes {
		if unicode.IsUpper(r) {
			serierunes[i] = unicode.ToLower(r)
		}
	}
	return string(serierunes)
}
