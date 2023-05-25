/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package match

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestMultiMatcher_GivenDefinition_MatchesExpectedStrings(t *testing.T) {
	t.Parallel()

	type expectation struct {
		value       string
		matchResult Result
	}

	cases := []struct {
		name       string
		definition string
		expected   []expectation
	}{
		{
			"Literal matches match all literals",
			"foo;bar;baz",
			[]expectation{
				{"foo", matchFound("foo")},
				{"bar", matchFound("bar")},
				{"baz", matchFound("baz")},
				{"zoo", matchNotFound()},
				{"Foo", matchFound("foo")},
				{"Bar", matchFound("bar")},
				{"Baz", matchFound("baz")},
				{"Zoo", matchNotFound()},
			},
		},
		{
			"Wildcard matches match different things",
			"f*;Ba*",
			[]expectation{
				{"foo", matchFound("f*")},
				{"bar", matchFound("Ba*")},
				{"baz", matchFound("Ba*")},
				{"zoo", matchNotFound()},
				{"Foo", matchFound("f*")},
				{"Bar", matchFound("Ba*")},
				{"Baz", matchFound("Ba*")},
				{"Zoo", matchNotFound()},
			},
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			m := newMultiMatcher(c.definition)
			for _, e := range c.expected {
				matchResult := m.Matches(e.value)
				g.Expect(matchResult).To(Equal(e.matchResult))
			}
		})
	}
}

func TestMultiMatcher_DoesNotShortCircuit(t *testing.T) {
	t.Parallel()

	/*
	 * For nested matchers to correctly offer suggestions, the MultiMatcher may not short circuit evaluations.
	 * Even once a match is found, later matchers must still be consulted so that they can accumulate candidates
	 * to use for diagnostic generation.
	 *
	 * To ensure no short circuiting occurs, we create some multimatchers where every branch will match a given
	 * probe value, and check that all branches were matched.
	 */

	cases := []struct {
		name          string
		definition    string
		probe         string
		expectedMatch string
	}{
		{"Similar literals", "foo;Foo", "foo", "foo"},
		{"Similar globs", "T*;*T", "that", "T*"},
		{"Mixed types", "B*;Ba?;baz", "baz", "B*"},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			m := newMultiMatcher(c.definition)

			matchResult := m.Matches(c.probe)
			g.Expect(matchResult.Matched).To(BeTrue())
			g.Expect(matchResult.MatchingPattern).To(Equal(c.expectedMatch))

			castMatches := m.(*multiMatcher)
			for _, nested := range castMatches.matchers {
				g.Expect(nested.WasMatched()).To(BeNil())
			}
		})
	}
}

func TestMultiMatcher_IsRestrictive_GivesExpectedResults(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name        string
		matcher     string
		restrictive bool
	}{
		{"Two unrestrictive matchers are unrestrictive", "*;", false},
		{"Restrictive wildcard is restrictive", "Foo*;", true},
		{"Restrictive literal is restrictive", "*;Foo", true},
	}

	for _, c := range cases {
		c := c
		t.Run(
			c.name,
			func(t *testing.T) {
				t.Parallel()
				g := NewGomegaWithT(t)
				matcher := newMultiMatcher(c.matcher)

				g.Expect(matcher.IsRestrictive()).To(Equal(c.restrictive))
			})
	}
}
