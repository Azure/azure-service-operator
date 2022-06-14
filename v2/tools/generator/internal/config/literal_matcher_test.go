/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package config

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestLiteralMatcher_Matches_GivesExpectedResults(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name     string
		literal  string
		value    string
		expected bool
	}{
		{"Exact match", "Foo", "Foo", true},
		{"Uppercase literal matches lowercase value", "FOO", "foo", true},
		{"Lowercase literal matches uppercase value", "foo", "FOO", true},
		{"Mixed case literal matches different mixed case value", "Foo", "foO", true},
		{"Different strings do not match", "Foo", "Bar", false},
		{"Leading whitespace in matcher matches", "  Foo", "Foo", true},
		{"Trailing whitespace in matcher matches", "Foo  ", "Foo", true},
		{"Leading whitespace in value matches", "Foo", "  Foo", true},
		{"Trailing whitespace in value matches", "Foo", "Foo  ", true},
	}

	for _, c := range cases {
		c := c
		t.Run(
			c.name,
			func(t *testing.T) {
				t.Parallel()
				g := NewGomegaWithT(t)
				matcher := newLiteralMatcher(c.literal)
				g.Expect(matcher.Matches(c.value)).To(Equal(c.expected))
			})
	}
}

func TestLiteralMatcher_IsRestrictive_GivesExpectedResults(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name        string
		literal     string
		restrictive bool
	}{
		{"Empty is not restrictive", "", false},
		{"Whitespace is not restrictive", "    ", false},
		{"Content is restrictive", "content", true},
	}

	for _, c := range cases {
		c := c
		t.Run(
			c.name,
			func(t *testing.T) {
				t.Parallel()
				g := NewGomegaWithT(t)
				matcher := newLiteralMatcher(c.literal)
				g.Expect(matcher.IsRestrictive()).To(Equal(c.restrictive))
			})
	}
}
