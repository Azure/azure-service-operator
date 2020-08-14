/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"go/ast"
	"regexp"
	"strings"
)

// Utility methods for adding comments

func addDocComments(commentList *[]*ast.Comment, comments []string, width int) {
	for _, comment := range comments {
		// Skip empty comments
		if comment == "" {
			continue
		}

		addDocComment(commentList, comment, width)
	}
}

func addDocComment(commentList *[]*ast.Comment, comment string, width int) {
	for _, c := range formatDocComment(comment, width) {
		line := strings.TrimSpace(c)

		if !strings.HasPrefix(line, "//") {
			line = "//" + line
		}

		if *commentList == nil {
			line = "\n" + line
		}

		*commentList = append(*commentList, &ast.Comment{
			Text: line,
		})
	}
}

// formatDocComment splits the supplied comment string up ready for use as a documentation comment
func formatDocComment(comment string, width int) []string {
	// Remove markdown bolding
	text := strings.ReplaceAll(comment, "**", "")

	// Turn <br> and <br/> into \n
	text = brRegex.ReplaceAllLiteralString(text, "\n")

	// Split into individual lines
	lines := strings.Split(text, "\n")

	// Trim whitespace
	for i, l := range lines {
		lines[i] = strings.TrimSpace(l)
	}

	// Wordwrap and return
	return docCommentWrap(lines, width)
}

var brRegex = regexp.MustCompile("<br[^/>]*/?>")

func docCommentWrap(lines []string, width int) []string {
	var result []string
	for _, l := range lines {
		result = append(result, wordWrap(l, width)...)
	}

	return result
}

func wordWrap(text string, width int) []string {
	var result []string

	start := 0
	for start < len(text) {
		finish := findBreakPoint(text, start, width)
		result = append(result, text[start:finish+1])
		start = finish + 1
	}

	return result
}

// findBreakPoint finds the character at which to break two lines
// Returned index points to the last character that should be included on the line
// If breaking at a space, this will give a trailing space, but allows for
// breaking at other points too as no characters will be omitted.
func findBreakPoint(line string, start int, width int) int {
	limit := start + width + 1
	if limit >= len(line) {
		return len(line) - 1
	}

	// Look for a word break within the line
	index := strings.LastIndex(line[start:limit], " ")
	if index >= 0 {
		return start + index
	}

	// Line contains continuous text, we don't want to break it in two, so find the end of it
	index = strings.Index(line[limit:], " ")
	if index >= 0 {
		return limit + index
	}

	return len(line) - 1
}
