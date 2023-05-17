/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package match

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/pkg/errors"

	"github.com/Azure/azure-service-operator/v2/internal/set"
)

// A globMatcher is used to match a string against a literal or wildcard
type globMatcher struct {
	glob       string         // The glob we're matching, may contain wildcards * and ?
	regex      *regexp.Regexp // A regular expression to match the glob
	matched    bool
	candidates set.Set[string]
}

var _ StringMatcher = &globMatcher{}

// newGlobMatcher returns a new matcher for handling wildcards
func newGlobMatcher(glob string) (StringMatcher, error) {
	if !HasWildCards(glob) {
		msg := fmt.Sprintf("glob string %q has no wildcards", glob)
		panic(msg)
	}

	regex, err := newGlobRegex(glob)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to create regexp for glob %q", glob)
	}

	return &globMatcher{
		glob:       glob,
		regex:      regex,
		candidates: make(set.Set[string]),
	}, nil
}

func (gm *globMatcher) Matches(term string) bool {
	if gm.regex.MatchString(term) {
		if !gm.matched {
			// First time we match, clear out our candidates as we won't be needing them
			gm.matched = true
			gm.candidates.Clear()
		}

		return true
	}

	if !gm.matched {
		// Still collecting candidates
		gm.candidates.Add(term)
	}

	return false
}

func (gm *globMatcher) WasMatched() error {
	if gm.matched {
		return nil
	}

	choices := set.AsSortedSlice(gm.candidates)
	return errors.Errorf(
		"no match for %q (available candidates were %s)",
		gm.glob,
		strings.Join(choices, ", "))
}

// IsRestrictive returns false if we are blank or a universal wildcard, true otherwise.
func (gm *globMatcher) IsRestrictive() bool {
	return gm.glob != "" && gm.glob != "*"
}

// String returns the literal we match
func (gm *globMatcher) String() string {
	return gm.glob
}

func newGlobRegex(glob string) (*regexp.Regexp, error) {
	g := regexp.QuoteMeta(glob)
	g = strings.ReplaceAll(g, "\\*", ".*")
	g = strings.ReplaceAll(g, "\\?", ".")
	g = "(?i)(^" + g + "$)"

	return regexp.Compile(g)
}

// HasWildCards returns true if the passed matcher string contains a wildcard, false otherwise.
func HasWildCards(matcher string) bool {
	return strings.ContainsRune(matcher, '*') || strings.ContainsRune(matcher, '?')
}
