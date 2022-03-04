/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package resolver

import (
	"fmt"
	"io"

	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/types"
)

type causer interface {
	error
	Cause() error
}

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
var _ causer = &ReferenceNotFound{}

func (e *ReferenceNotFound) Error() string {
	return fmt.Sprintf("%s does not exist (%s)", e.NamespacedName, e.cause)
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

func (e *ReferenceNotFound) Format(s fmt.State, verb rune) {
	format(e, s, verb)
}

type SecretNotFound struct {
	NamespacedName types.NamespacedName
	cause          error
}

func NewSecretNotFoundError(name types.NamespacedName, cause error) *SecretNotFound {
	return &SecretNotFound{
		NamespacedName: name,
		cause:          cause,
	}
}

var _ error = &SecretNotFound{}
var _ causer = &SecretNotFound{}

func (e *SecretNotFound) Error() string {
	return fmt.Sprintf("%s does not exist (%s)", e.NamespacedName, e.cause)
}

func (e *SecretNotFound) Is(err error) bool {
	var typedErr *SecretNotFound
	if errors.As(err, &typedErr) {
		return e.NamespacedName == typedErr.NamespacedName
	}
	return false
}

func (e *SecretNotFound) Cause() error {
	return e.cause
}

func (e *SecretNotFound) Format(s fmt.State, verb rune) {
	format(e, s, verb)
}

// This was adapted from the function in errors
func format(e causer, s fmt.State, verb rune) {
	switch verb {
	case 'v':
		if s.Flag('+') {
			n, _ := fmt.Fprintf(s, "%s", e.Cause())
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
