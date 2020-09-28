/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package armresourceresolver

import (
	"fmt"
	"io"

	"k8s.io/apimachinery/pkg/types"
)

type OwnerNotFound struct {
	OwnerName types.NamespacedName
	cause     error
}

func NewOwnerNotFoundError(ownerName types.NamespacedName, cause error) *OwnerNotFound {
	return &OwnerNotFound{
		OwnerName: ownerName,
		cause:     cause,
	}
}

var _ error = &OwnerNotFound{}

func (e *OwnerNotFound) Error() string {
	return fmt.Sprintf("%s does not exist", e.OwnerName)
}

func (e *OwnerNotFound) Is(err error) bool {
	if typedErr, ok := err.(*OwnerNotFound); ok {
		return e.OwnerName == typedErr.OwnerName
	}
	return false
}

func (e *OwnerNotFound) Cause() error {
	return e.cause
}

// This was adapted from the function in errors
func (e *OwnerNotFound) Format(s fmt.State, verb rune) {
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
