package shared

import "strings"

func IsNotFalse(s string) bool {
	return !IsFalse(s)
}

func IsFalse(s string) bool {
	return strings.ToLower(s) == "false" || strings.HasPrefix(strings.ToLower(s), "n") || strings.ToLower(s) == "0"
}
