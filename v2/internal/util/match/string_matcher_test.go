/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package match

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestStringMatcher_GivenMatcher_ReturnsExpectedResults(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name     string
		matcher  string
		value    string
		expected bool
	}{
		{"Case sensitive literal match", "Foo", "Foo", true},
		{"Case insensitive literal match", "FOO", "foo", true},
		{"Different strings do not match", "Foo", "Bar", false},
		{"Simple wildcard matches", "*", "Baz", true},
		{"Prefix with wildcard matches", "F*", "Foo", true},
		{"Suffix with wildcard matches", "*z", "Baz", true},
		{"Multiple wildcards", "F*;BA*", "Foo", true},
		{"Multiple wildcards, no match", "F*;BA*", "Zoo", false},
	}

	for _, c := range cases {
		c := c
		t.Run(
			c.name,
			func(t *testing.T) {
				t.Parallel()
				g := NewGomegaWithT(t)
				matcher, err := NewStringMatcher(c.matcher)
				g.Expect(err).ToNot(HaveOccurred())

				g.Expect(matcher.Matches(c.value)).To(Equal(c.expected))
			})
	}
}
