/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package config

import (
	"strings"
)

// literalMatcher is a StringMatcher that provides a case-insensitive match against a given string
type literalMatcher struct {
	literal string
	matched bool
	advisor *TypoAdvisor
}

var _ StringMatcher = &literalMatcher{}

// newLiteralMatcher returns a new literalMatcher for the given string
func newLiteralMatcher(literal string) *literalMatcher {
	return &literalMatcher{
		literal: literal,
		advisor: NewTypoAdvisor(),
	}
}

// Matches returns true if the passed value is a case-insensitive match with our configured literal
func (lm *literalMatcher) Matches(value string) bool {
	if strings.EqualFold(lm.literal, value) {
		if !lm.matched {
			// First time we match, clear out our advisory as we won't be using it
			lm.matched = true
			lm.advisor.ClearTerms()
		}

		return true
	}

	if !lm.matched {
		// Still collecting potential suggestions
		lm.advisor.AddTerm(value)
	}

	return false
}

// WasMatched returns an error if we didn't match anything, nil otherwise
func (lm *literalMatcher) WasMatched() error {
	if lm.matched {
		return nil
	}

	return lm.advisor.Errorf(lm.literal, "no match for %q", lm.literal)
}

// String returns the literal we match
func (lm *literalMatcher) String() string {
	return lm.literal
}
