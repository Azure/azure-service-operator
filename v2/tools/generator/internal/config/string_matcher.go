/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package config

import (
	"fmt"
)

// StringMatcher is an interface implemented by predicates used to test string values
type StringMatcher interface {
	fmt.Stringer
	// Matches returns true if the provided value is matched by the matcher
	Matches(value string) bool
	// WasMatched returns nil if the matcher had a match, otherwise returning a diagnostic error
	WasMatched() error
}
