/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package config

import (
	"errors"
	"fmt"
	"sort"
	"strings"
)

// NotConfiguredError is returned when requested configuration is not found
type NotConfiguredError struct {
	message    string
	optionKind string
	options    []string
}

// Ensure we implement the error interface
var _ error = NotConfiguredError{}

func NewNotConfiguredError(message string) NotConfiguredError {
	return NotConfiguredError{
		message: message,
	}
}

// IsNotConfiguredError returns true if the passed error is a NotConfiguredError, false otherwise
func IsNotConfiguredError(err error) bool {
	var nce NotConfiguredError
	return errors.As(err, &nce)
}

// WithOptions configures our error to include a sequence of available options that are available.
// Including this information in the message helps end users identify when there is a spelling or other error.
func (e NotConfiguredError) WithOptions(kind string, options []string) NotConfiguredError {
	e.optionKind = kind
	e.options = options
	return e
}

// Return a string representing the error
func (e NotConfiguredError) Error() string {
	if len(e.options) == 0 {
		return e.message
	}

	sort.Strings(e.options)
	return fmt.Sprintf(
		"%s (%d available %s are %s)",
		e.message,
		len(e.options),
		e.optionKind,
		strings.Join(e.options, "; "))
}
