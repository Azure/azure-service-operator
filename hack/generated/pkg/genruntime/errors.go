/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package genruntime

import (
	"fmt"
	"io"

	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/types"
)

type ReferenceNotFound struct {
	NamespacedName types.NamespacedName
	cause          error
}

func NewReferenceNotFoundError(name types.NamespacedName, cause error) *ReferenceNotFound {
	return &ReferenceNotFound{
		NamespacedName: name,
		cause:          cause,
	}
}

var _ error = &ReferenceNotFound{}

func (e *ReferenceNotFound) Error() string {
	return fmt.Sprintf("%s does not exist", e.NamespacedName)
}

func (e *ReferenceNotFound) Is(err error) bool {
	var typedErr *ReferenceNotFound
	if errors.As(err, &typedErr) {
		return e.NamespacedName == typedErr.NamespacedName
	}
	return false
}

func (e *ReferenceNotFound) Cause() error {
	return e.cause
}

// This was adapted from the function in errors
func (e *ReferenceNotFound) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		if s.Flag('+') {
			n, _ := fmt.Fprintf(s, "%+v", e.Cause())
			if n > 0 {
				_, _ = fmt.Fprintf(s, "\n")
			}
			_, _ = io.WriteString(s, e.Error())
			return
		}
		fallthrough
	case 's':
		_, _ = io.WriteString(s, e.Error())
	case 'q':
		_, _ = fmt.Fprintf(s, "%q", e.Error())
	}
}
