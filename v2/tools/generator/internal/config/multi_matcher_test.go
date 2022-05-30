/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package config

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestMultiMatcher_GivenDefinition_MatchesExpectedStrings(t *testing.T) {
	t.Parallel()

	type expectation struct {
		value string
		match bool
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
				{"foo", true},
				{"bar", true},
				{"baz", true},
				{"zoo", false},
				{"Foo", true},
				{"Bar", true},
				{"Baz", true},
				{"Zoo", false},
			},
		},
		{
			"Wildcard matches match different things",
			"f*;Ba*",
			[]expectation{
				{"foo", true},
				{"bar", true},
				{"baz", true},
				{"zoo", false},
				{"Foo", true},
				{"Bar", true},
				{"Baz", true},
				{"Zoo", false},
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
				g.Expect(m.Matches(e.value)).To(Equal(e.match))
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
		name       string
		definition string
		probe      string
	}{
		{"Similar literals", "foo;Foo", "foo"},
		{"Similar globs", "T*;*T", "that"},
		{"Mixed types", "B*;Ba?;baz", "baz"},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			m := newMultiMatcher(c.definition)
			g.Expect(m.Matches(c.probe)).To(BeTrue())
			for _, nested := range m.matchers {
				g.Expect(nested.WasMatched()).To(BeNil())
			}
		})
	}

}
