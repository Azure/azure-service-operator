/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"regexp"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/samber/lo"
)

func TestStringValidations_MergeWith_DeduplicatesPatterns(t *testing.T) {
	t.Parallel()

	pattern1 := regexp.MustCompile("^[a-z]+$")
	pattern2 := regexp.MustCompile("^[0-9]+$")
	pattern3 := regexp.MustCompile("^[A-Z]+$")
	// Same pattern as pattern1 but different pointer
	pattern1Duplicate := regexp.MustCompile("^[a-z]+$")

	cases := []struct {
		name             string
		svPatterns       []*regexp.Regexp
		otherPatterns    []*regexp.Regexp
		expectedPatterns []string // Compare by string representation
	}{
		{
			name:             "Empty patterns merged with empty patterns",
			svPatterns:       nil,
			otherPatterns:    nil,
			expectedPatterns: nil,
		},
		{
			name:             "Non-empty merged with empty",
			svPatterns:       []*regexp.Regexp{pattern1, pattern2},
			otherPatterns:    nil,
			expectedPatterns: []string{"^[a-z]+$", "^[0-9]+$"},
		},
		{
			name:             "Empty merged with non-empty",
			svPatterns:       nil,
			otherPatterns:    []*regexp.Regexp{pattern1, pattern2},
			expectedPatterns: []string{"^[a-z]+$", "^[0-9]+$"},
		},
		{
			name:             "Different patterns are all kept",
			svPatterns:       []*regexp.Regexp{pattern1},
			otherPatterns:    []*regexp.Regexp{pattern2, pattern3},
			expectedPatterns: []string{"^[a-z]+$", "^[0-9]+$", "^[A-Z]+$"},
		},
		{
			name:             "Same pointer pattern is deduplicated",
			svPatterns:       []*regexp.Regexp{pattern1},
			otherPatterns:    []*regexp.Regexp{pattern1},
			expectedPatterns: []string{"^[a-z]+$"},
		},
		{
			name:             "Same string pattern different pointers is deduplicated",
			svPatterns:       []*regexp.Regexp{pattern1},
			otherPatterns:    []*regexp.Regexp{pattern1Duplicate},
			expectedPatterns: []string{"^[a-z]+$"},
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			sv := StringValidations{Patterns: c.svPatterns}
			other := StringValidations{Patterns: c.otherPatterns}

			merged, err := sv.MergeWith(other)
			g.Expect(err).ToNot(HaveOccurred())

			mergedSV, ok := merged.(StringValidations)
			g.Expect(ok).To(BeTrue())

			// Convert result patterns to strings for comparison
			var resultPatterns []string
			if mergedSV.Patterns != nil {
				resultPatterns = lo.Map(mergedSV.Patterns, func(p *regexp.Regexp, _ int) string {
					return p.String()
				})
			}

			g.Expect(resultPatterns).To(Equal(c.expectedPatterns))
		})
	}
}

func TestAppendUniquePatterns_PreservesOrder(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	pattern1 := regexp.MustCompile("^a$")
	pattern2 := regexp.MustCompile("^b$")
	pattern3 := regexp.MustCompile("^c$")
	pattern4 := regexp.MustCompile("^d$")

	existing := []*regexp.Regexp{pattern1, pattern2}
	new := []*regexp.Regexp{pattern3, pattern4}

	result := appendUniquePatterns(existing, new)

	// Verify order is preserved: existing patterns first, then new patterns
	g.Expect(result).To(HaveLen(4))
	g.Expect(result[0].String()).To(Equal("^a$"))
	g.Expect(result[1].String()).To(Equal("^b$"))
	g.Expect(result[2].String()).To(Equal("^c$"))
	g.Expect(result[3].String()).To(Equal("^d$"))
}
