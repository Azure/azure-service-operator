/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package importresources

import (
	"fmt"

	"k8s.io/apimachinery/pkg/runtime/schema"
)

type ImportError struct {
	err  error            // The error that caused the import to fail
	gk   schema.GroupKind // The GroupKind of the resource that failed to import
	name string           // The name of the resource that failed to import
}

var _ error = &ImportError{}

func MakeImportError(err error, gk schema.GroupKind, name string) ImportError {
	return ImportError{
		err:  err,
		gk:   gk,
		name: name,
	}
}

func (ie *ImportError) Error() string {
	return fmt.Sprintf("failed to import %s %s: %s", ie.gk, ie.name, ie.err.Error())
}

func (ie *ImportError) Unwrap() error {
	return ie.err
}
