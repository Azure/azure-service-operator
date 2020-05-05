/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package jsonast

import (
	"testing"

	. "github.com/onsi/gomega"
)

func Test_CreateGlobbingRegex_ReturnsExpectedRegex(t *testing.T) {

	cases := []struct {
		glob  string
		regex string
	}{
		{"*preview", "^.*preview$"},
		{"*.bak", "^.*\\.bak$"},
		{"2014*", "^2014.*$"},
		{"2014-??-??", "^2014-..-..$"},
	}

	for _, c := range cases {
		c := c
		t.Run(c.glob, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)
			r := createGlobbingRegex(c.glob)
			s := r.String()
			g.Expect(s).To(Equal(c.regex))
		})
	}
}

func Test_GlobbingRegex_MatchesExpectedStrings(t *testing.T) {

	cases := []struct {
		regex       string
		candidate   string
		shouldMatch bool
	}{
		{"*preview", "2020-02-01preview", true},
		{"*preview", "2020-02-01", false},
		{"2020-*", "2020-02-01", true},
		{"2020-*", "2020-01-01", true},
		{"2020-*", "2019-01-01", false},
		{"2020-*", "2015-07-01", false},
	}

	for _, c := range cases {
		c := c
		t.Run(c.candidate, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)
			r := createGlobbingRegex(c.regex)
			match := r.MatchString(c.candidate)
			g.Expect(match).To(Equal(c.shouldMatch))
		})
	}
}
