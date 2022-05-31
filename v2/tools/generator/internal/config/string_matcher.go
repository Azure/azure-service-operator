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

// NewStringMatcher returns a matcher for the specified string
// Different strings may return different implementations:
// o If the string contains '*' or '?' a globbing wildcard matcher
// o Otherwise a case-insensitive literal string matcher
func NewStringMatcher(matcher string) StringMatcher {
	if matcher == "" {
		return newAlwaysMatcher()
	}

	if HasMultipleMatchers(matcher) {
		return newMultiMatcher(matcher)
	}

	if HasWildCards(matcher) {
		return newGlobMatcher(matcher)
	}

	return newLiteralMatcher(matcher)
}
