/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astbuilder

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestDocumentationCommentFormatting(t *testing.T) {
	t.Parallel()

	cases := []struct {
		comment string
		results []string
	}{
		// Expect short single line comments to be unchanged
		{"foo", []string{"foo"}},
		{"bar", []string{"bar"}},
		{"baz", []string{"baz"}},
		// Leading and trailing whitespace is trimmed
		{"    foo", []string{"foo"}},
		{"foo    ", []string{"foo"}},
		{"  foo  ", []string{"foo"}},
		// Expect comments with embedded newlines to be split
		{"foo\nbar", []string{"foo", "bar"}},
		{"foo\nbar\nbaz", []string{"foo", "bar", "baz"}},
		// Expect comments with html style <br> to be split
		{"foo<br>bar", []string{"foo", "bar"}},
		{"foo<br>bar<br>baz", []string{"foo", "bar", "baz"}},
		{"foo<br/>bar", []string{"foo", "bar"}},
		{"foo<br/>bar<br/>baz", []string{"foo", "bar", "baz"}},
		// Expect markdown bold to be removed
		{"**foo**\nbar", []string{"foo", "bar"}},
		{"foo\n**bar**", []string{"foo", "bar"}},
		{"foo\n**bar**\nbaz", []string{"foo", "bar", "baz"}},
		{"", []string{}},
		// Expect long lines to be wrapped
		{
			"Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.",
			[]string{
				"Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do ",
				"eiusmod tempor incididunt ut labore et dolore magna aliqua.",
			}},
	}

	for _, c := range cases {
		if c.comment == "" {
			continue
		}
		c := c
		t.Run(c.comment, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)
			lines := formatComment(c.comment, 64)
			g.Expect(lines).To(Equal(c.results))
		})
	}
}

func TestWordWrap(t *testing.T) {
	t.Parallel()

	cases := []struct {
		text    string
		width   int
		results []string
	}{
		{"this is a simple line of text", 15, []string{"this is a ", "simple line of ", "text"}},
		{"this is a simple line of text", 16, []string{"this is a simple ", "line of text"}},
		{"this is a simple line of text", 20, []string{"this is a simple ", "line of text"}},
		{"this is a simple line of text", 21, []string{"this is a simple line ", "of text"}},
		{"", 0, []string{}},
	}

	for _, c := range cases {
		if c.width == 0 {
			continue
		}
		c := c
		t.Run(c.text, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)
			lines := WordWrap(c.text, c.width)
			g.Expect(lines).To(Equal(c.results))
		})
	}
}
