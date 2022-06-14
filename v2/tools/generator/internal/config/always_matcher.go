/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package config

// alwaysMatcher is a matcher that always matches
type alwaysMatcher struct{}

var _ StringMatcher = &alwaysMatcher{}

// newAlwaysMatcher returns a new instance of an alwaysMatcher
func newAlwaysMatcher() *alwaysMatcher {
	return &alwaysMatcher{}
}

// String returns a string representing the matcher
func (a alwaysMatcher) String() string {
	return ""
}

// Matches returns true as we always match
func (a alwaysMatcher) Matches(value string) bool {
	return true
}

// WasMatched returns nil because we always have matched
func (a alwaysMatcher) WasMatched() error {
	return nil
}

// IsRestrictive returns false because the always matcher doesn't restrict anything
func (a alwaysMatcher) IsRestrictive() bool {
	return false
}
