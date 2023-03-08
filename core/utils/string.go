package utils

import "strings"

func IsLowercaseCharacter(char int32) bool {
	return char >= 'a' && char <= 'z'
}

func IsUppercaseCharacter(char int32) bool {
	return char >= 'A' && char <= 'Z'
}

func ToUpperInitial(value string) string {
	return strings.ToUpper(value[:1]) + value[1:]
}
