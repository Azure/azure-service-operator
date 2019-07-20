package helpers

import (
	"fmt"
	"strings"

	uuid "github.com/satori/go.uuid"
)

func ContainsString(slice []string, s string) bool {
	for _, item := range slice {
		if item == s {
			return true
		}
	}
	return false
}

func RemoveString(slice []string, s string) (result []string) {
	for _, item := range slice {
		if item == s {
			continue
		}
		result = append(result, item)
	}
	return
}

func AzureResourceName(kind string) string {
	var name string
	switch kind {
	case "Storage":
		name = fmt.Sprintf("aso%s", strings.ReplaceAll(uuid.NewV4().String(), "-", ""))[:24]
	}
	return name
}
