package organizer

import (
	"strings"
	"unicode"
)

// GetAlphabetFolder retorna a letra maiúscula correspondente ou "others" para símbolos/números
func GetAlphabetFolder(filename string) string {
	if filename == "" {
		return "others"
	}

	firstChar := rune(strings.ToUpper(filename)[0])

	if unicode.IsLetter(firstChar) {
		return string(firstChar)
	}

	return "others"
}
