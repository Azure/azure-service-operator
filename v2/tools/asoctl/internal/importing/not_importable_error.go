/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package importing

import (
	"fmt"

	"k8s.io/apimachinery/pkg/runtime/schema"
)

// NotImportableError is an error that indicates that a resource cannot be imported for a reason we know about.
// This allows us to continue the import even if some expected errors occur.
type NotImportableError struct {
	GroupKind schema.GroupKind
	Name      string
	Because   string
}

// Ensure we implement the error interface
var _ error = &NotImportableError{}

func NewNotImportableError(groupKind schema.GroupKind, name string, because string) NotImportableError {
	return NotImportableError{
		GroupKind: groupKind,
		Name:      name,
		Because:   because,
	}
}

func (e NotImportableError) Error() string {
	return fmt.Sprintf(
		"%s/%s %s was skipped because %s",
		e.GroupKind.Group,
		e.GroupKind.Kind,
		e.Name,
		e.Because)
}
