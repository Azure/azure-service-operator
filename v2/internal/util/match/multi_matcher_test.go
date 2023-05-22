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
		value          string
		match          bool
		expectedDetail string
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
				{"foo", true, "foo"},
				{"bar", true, "bar"},
				{"baz", true, "baz"},
				{"zoo", false, ""},
				{"Foo", true, "foo"},
				{"Bar", true, "bar"},
				{"Baz", true, "baz"},
				{"Zoo", false, ""},
			},
		},
		{
			"Wildcard matches match different things",
			"f*;Ba*",
			[]expectation{
				{"foo", true, "f*"},
				{"bar", true, "Ba*"},
				{"baz", true, "Ba*"},
				{"zoo", false, ""},
				{"Foo", true, "f*"},
				{"Bar", true, "Ba*"},
				{"Baz", true, "Ba*"},
				{"Zoo", false, ""},
			},
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			m, err := newMultiMatcher(c.definition)
			g.Expect(err).ToNot(HaveOccurred())
			for _, e := range c.expected {
				matches, detail := m.MatchesDetailed(e.value)
				g.Expect(matches).To(Equal(e.match))
				g.Expect(detail).To(Equal(e.expectedDetail))
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

			m, err := newMultiMatcher(c.definition)
			g.Expect(err).ToNot(HaveOccurred())

			matches, detail := m.MatchesDetailed(c.probe)
			g.Expect(matches).To(BeTrue())
			g.Expect(detail).To(Equal(c.expectedMatch))

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
				matcher, err := newMultiMatcher(c.matcher)
				g.Expect(err).ToNot(HaveOccurred())

				g.Expect(matcher.IsRestrictive()).To(Equal(c.restrictive))
			})
	}
}
