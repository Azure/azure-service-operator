/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package core

import (
	"fmt"

	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/runtime/schema"
	kerrors "k8s.io/apimachinery/pkg/util/errors"
)

func AsTypedError[T error](err error) (T, bool) {
	var typedErr T
	if errors.As(err, &typedErr) {
		return typedErr, true
	}

	// Also deal with the possibility that this is a kerrors.Aggregate
	var aggregate kerrors.Aggregate
	if errors.As(err, &aggregate) {
		for _, e := range aggregate.Errors() {
			// This is a bit hacky but allows us to pick out the first error and raise on that
			if result, ok := AsTypedError[T](e); ok {
				return result, true
			}
		}
	}

	return typedErr, false
}

func AsNotOwnedError(err error) (*NotOwnedError, bool) {
	return AsTypedError[*NotOwnedError](err)
}

// NotOwnedError indicates the target resource is not owned by the resource attempting to write it
type NotOwnedError struct {
	Namespace  string
	TargetName string
	TargetType string
	SourceName string
	SourceType string
}

func NewNotOwnedError(namespace string, name string, gvk schema.GroupVersionKind, sourceName string, sourceGvk schema.GroupVersionKind) error {
	kindStr := gvk.GroupKind().String()
	sourceKindStr := sourceGvk.GroupKind().String()
	return &NotOwnedError{
		Namespace:  namespace,
		TargetName: name,
		TargetType: kindStr,
		SourceName: sourceName,
		SourceType: sourceKindStr,
	}
}

var _ error = &NotOwnedError{}

func (e *NotOwnedError) Error() string {
	return fmt.Sprintf("cannot overwrite %s %s/%s which is not owned by %s %s/%s",
		e.TargetType,
		e.Namespace,
		e.TargetName,
		e.SourceType,
		e.Namespace,
		e.SourceName)
}
