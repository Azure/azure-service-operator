package shared

import "strings"

func IsNotFalse(s string) bool {
	return strings.ToLower(s) != "false" && strings.ToLower(s) != "no"
}
