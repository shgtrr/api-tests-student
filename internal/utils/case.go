package utils

import (
	"strings"
	"unicode"
)

// mixcase переводит каждый второй символ строки в верхний регистр
func MixCase(text string) string {
	characters := []rune(strings.ToLower(text))
	for index := 0; index < len(characters); index += 2 {
		characters[index] = unicode.ToUpper(characters[index])
	}

	return string(characters)
}
