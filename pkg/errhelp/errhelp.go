package errhelp

import (
	"strings"
)

// IsParentNotFound checks if the error is about a parent resrouce not existing
func IsParentNotFound(err error) bool {
	return strings.Contains(err.Error(), "ParentResourceNotFound")
}

// IsGroupNotFound checks if error is about resource group not existing
func IsGroupNotFound(err error) bool {
	return strings.Contains(err.Error(), "ResourceGroupNotFound")
}

// IsNotActive checks if error is mentioning a non active resource
func IsNotActive(err error) bool {
	return strings.Contains(err.Error(), "not active")
}
