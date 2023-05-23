/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package match

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestGlobMatcher_GivenTerms_MatchesExpectedStrings(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name          string
		glob          string
		term          string
		expectedMatch bool
	}{
		{"Star wildcard matches short string", "*", "foo", true},
		{"Star wildcard matches longstring", "*", "foobarbaz", true},
		{"Question-mark wildcard matches letter", "foo?", "fool", true},
		{"Question-mark wildcard matches digit", "foo?", "foo7", true},
	}

	for _, c := range cases {
		c := c
		t.Run(
			c.name,
			func(t *testing.T) {
				t.Parallel()
				g := NewGomegaWithT(t)
				matcher := newGlobMatcher(c.glob)

				g.Expect(matcher.Matches(c.term).Matched).To(Equal(c.expectedMatch))
			})
	}
}

func TestGlobMatcher_IsRestrictive_GivesExpectedResults(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name        string
		glob        string
		restrictive bool
	}{
		{"Wildcard is not restrictive", "*", false},
		{"Complex match is restrictive", "Foo*", true},
	}

	for _, c := range cases {
		c := c
		t.Run(
			c.name,
			func(t *testing.T) {
				t.Parallel()
				g := NewGomegaWithT(t)
				matcher := newGlobMatcher(c.glob)

				g.Expect(matcher.IsRestrictive()).To(Equal(c.restrictive))
			})
	}
}
