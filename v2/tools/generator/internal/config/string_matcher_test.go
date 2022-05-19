/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package config

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
	}

	for _, c := range cases {
		c := c
		t.Run(
			c.name,
			func(t *testing.T) {
				t.Parallel()
				g := NewGomegaWithT(t)
				matcher := NewStringMatcher(c.matcher)
				g.Expect(matcher.Matches(c.value)).To(Equal(c.expected))
			})
	}

}
