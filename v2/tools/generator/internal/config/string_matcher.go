/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package config

import (
	"github.com/Azure/azure-service-operator/v2/internal/util/match"
)

// NewStringMatcher returns a matcher for the specified string
// Different strings may return different implementations:
// o If the string contains '*' or '?' a globbing wildcard matcher
// o Otherwise a case-insensitive literal string matcher
func NewStringMatcher(matcher string) match.StringMatcher {
	m := match.NewStringMatcher(matcher)

	return m
}
