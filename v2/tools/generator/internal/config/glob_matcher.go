package config

import (
	"fmt"
	"regexp"
	"strings"
	"sync"

	"github.com/Azure/azure-service-operator/v2/internal/set"
	"github.com/pkg/errors"
)

// A globMatcher is used to match a string against a literal or wildcard
type globMatcher struct {
	glob       string         // The glob we're matching, may contain wildcards * and ?
	regex      *regexp.Regexp // A regular expression to match the glob
	once       sync.Once
	matched    bool
	candidates set.Set[string]
}

var _ StringMatcher = &globMatcher{}

// newGlobMatcher returns a new matcher for handling wildcards
func newGlobMatcher(glob string) *globMatcher {
	// Guard against misuse
	if !HasWildCards(glob) {
		msg := fmt.Sprintf("glob string %q has no wildcards", glob)
		panic(msg)
	}

	return &globMatcher{
		glob:       glob,
		candidates: make(set.Set[string]),
	}
}

func (gm *globMatcher) Matches(term string) bool {
	gm.once.Do(gm.createRegex)

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

func (gm *globMatcher) createRegex() {
	g := regexp.QuoteMeta(gm.glob)
	g = strings.ReplaceAll(g, "\\*", ".*")
	g = strings.ReplaceAll(g, "\\?", ".")
	g = "(?i)(^" + g + "$)"
	gm.regex = regexp.MustCompile(g)
}

// HasWildCards returns true if the passed matcher string contains a wildcard, false otherwise.
func HasWildCards(matcher string) bool {
	return strings.ContainsRune(matcher, '*') || strings.ContainsRune(matcher, '?')
}

// String returns the literal we match
func (gm *globMatcher) String() string {
	return gm.glob
}
