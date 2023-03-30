/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package importing

import (
	"fmt"
)

// NotImportableError is an error that indicates that a resource cannot be imported for a reason we know about.
// This allows us to continue the import even if some expected errors occur.
type NotImportableError struct {
	ResourceType string
	Name         string
	Because      string
}

// Ensure we implement the error interface
var _ error = NotImportableError{}

func NewNotImportableError(resourceType string, name string, because string) NotImportableError {
	return NotImportableError{
		ResourceType: resourceType,
		Name:         name,
		Because:      because,
	}
}

func (e NotImportableError) Error() string {
	return fmt.Sprintf("resource %s/%s is not importable because %s", e.ResourceType, e.Name, e.Because)
}
