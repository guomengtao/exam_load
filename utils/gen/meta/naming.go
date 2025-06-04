package meta

import (
	"strings"
	"unicode"
)

// CamelCase converts a snake_case string to CamelCase.
// Example: "user_name" -> "UserName"
func CamelCase(s string) string {
	parts := strings.Split(s, "_")
	for i, part := range parts {
		if part == "" {
			continue
		}
		runes := []rune(part)
		runes[0] = unicode.ToUpper(runes[0])
		for j := 1; j < len(runes); j++ {
			runes[j] = unicode.ToLower(runes[j])
		}
		parts[i] = string(runes)
	}
	return strings.Join(parts, "")
}

// LowerCamelCase converts a snake_case string to lowerCamelCase.
// Example: "user_name" -> "userName"
func LowerCamelCase(s string) string {
	camel := CamelCase(s)
	if len(camel) == 0 {
		return ""
	}
	runes := []rune(camel)
	runes[0] = unicode.ToLower(runes[0])
	return string(runes)
}

// SnakeCase converts a CamelCase or lowerCamelCase string to snake_case.
// Example: "UserName" -> "user_name", "userName" -> "user_name"
func SnakeCase(s string) string {
	var result []rune
	for i, r := range s {
		if unicode.IsUpper(r) {
			if i > 0 {
				result = append(result, '_')
			}
			result = append(result, unicode.ToLower(r))
		} else {
			result = append(result, r)
		}
	}
	return string(result)
}
