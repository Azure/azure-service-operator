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
