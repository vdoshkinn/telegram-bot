package common

import "strings"

var charactersToEscape = []string{"_", "*", "[", "]", "(", ")", "~", "`", ">", "#", "+", "-", "=", "|", "{", "}", ".", "!"}

func EscapeCharacters(s string) string {
	for _, char := range charactersToEscape {
		s = strings.ReplaceAll(s, char, "\\"+char)
	}
	return s
}
