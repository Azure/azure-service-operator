/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package importing

import (
	"fmt"

	"k8s.io/apimachinery/pkg/runtime/schema"
)

// ImportSkippedError is an error that indicates that a resource cannot be imported for a reason we know about.
// This allows us to continue the import even if some expected errors occur.
type ImportSkippedError struct {
	GroupKind schema.GroupKind
	Name      string
	Because   string
	Resource  ImportableResource
}

// Ensure we implement the error interface
var _ error = &ImportSkippedError{}

func NewImportSkippedError(
	groupKind schema.GroupKind,
	name string,
	because string,
	resource ImportableResource,
) *ImportSkippedError {
	return &ImportSkippedError{
		GroupKind: groupKind,
		Name:      name,
		Because:   because,
		Resource:  resource,
	}
}

func (e ImportSkippedError) Error() string {
	return fmt.Sprintf(
		"%s/%s %s was skipped because %s",
		e.GroupKind.Group,
		e.GroupKind.Kind,
		e.Name,
		e.Because)
}
