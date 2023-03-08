package codestyle

import (
	"github.com/Haato3o/eskema/core/utils"
	"strings"
)

func ToSnakeCase(value string) string {
	normalized := normalizeNamingStyle(value)

	return strings.ReplaceAll(normalized, "+", "_")
}

func ToCamelCase(value string) string {
	normalized := normalizeNamingStyle(value)
	tokens := strings.Split(normalized, "+")

	var builder strings.Builder

	for i, token := range tokens {
		if i == 0 {
			builder.WriteString(token)
		} else {
			builder.WriteString(utils.ToUpperInitial(token))
		}
	}

	return builder.String()
}

func ToPascalCase(value string) string {
	camelCase := ToCamelCase(value)

	return utils.ToUpperInitial(camelCase)
}

func normalizeNamingStyle(value string) string {
	if isSnakeCase(value) {
		return normalizeFromSnakeCase(value)
	}

	return normalizeFromPascalCase(value)
}

func normalizeFromPascalCase(value string) string {
	values := make([]string, 0)

	var builder strings.Builder
	for i, char := range value {
		isUppercase := utils.IsUppercaseCharacter(char)

		if i != 0 && isUppercase {
			values = append(values, builder.String())
			builder.Reset()
		}

		currentCharacter := strings.ToLower(string(char))
		builder.WriteString(currentCharacter)
	}

	if builder.Len() > 0 {
		values = append(values, builder.String())
	}

	return strings.Join(values, "+")
}

func isSnakeCase(value string) bool {
	if length := len(strings.Split(value, "_")); length > 1 {
		return true
	}

	for _, char := range value {
		if utils.IsLowercaseCharacter(char) {
			return false
		}
	}

	return true
}

func normalizeFromSnakeCase(value string) string {
	values := strings.Split(strings.ToLower(value), "_")

	return strings.Join(values, "+")
}
