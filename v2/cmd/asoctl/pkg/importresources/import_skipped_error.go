/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package importresources

import (
	"fmt"

	"k8s.io/apimachinery/pkg/runtime/schema"
)

// SkippedError is an error that indicates that a resource cannot be imported for a reason we know about.
// This allows us to continue the import even if some expected errors occur.
type SkippedError struct {
	GroupKind schema.GroupKind
	Name      string
	Because   string
	Resource  ImportableResource
}

// Ensure we implement the error interface
var _ error = &SkippedError{}

func NewSkippedError(
	groupKind schema.GroupKind,
	name string,
	because string,
	resource ImportableResource,
) *SkippedError {
	return &SkippedError{
		GroupKind: groupKind,
		Name:      name,
		Because:   because,
		Resource:  resource,
	}
}

func (e SkippedError) Error() string {
	return fmt.Sprintf(
		"%s/%s %s was skipped because %s",
		e.GroupKind.Group,
		e.GroupKind.Kind,
		e.Name,
		e.Because)
}
