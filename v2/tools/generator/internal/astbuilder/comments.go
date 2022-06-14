/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astbuilder

import (
	"regexp"
	"strings"

	"github.com/dave/dst"
)

const DefaultCommentWrapWidth = 120

// AddWrappedComments adds comments to the specified list, wrapping text to the specified width as
// it goes. Respects any existing line breaks specified by \n or <br>
func AddWrappedComments(commentList *dst.Decorations, comments []string) {
	for _, comment := range comments {
		// Skip empty comments
		if comment == "" {
			continue
		}

		AddWrappedComment(commentList, comment)
	}
}

// AddWrappedComment adds a single comment to the specified list, wrapping text as it goes.
// Respects any existing line breaks specified by \n or <br>
func AddWrappedComment(commentList *dst.Decorations, comment string) {
	AddWrappedCommentAtWidth(commentList, comment, DefaultCommentWrapWidth)
}

// AddWrappedCommentAtWidth adds a single comment to the specified list, wrapping text to the specified
// width as it goes. Respects any existing line breaks specified by \n or <br>
func AddWrappedCommentAtWidth(commentList *dst.Decorations, comment string, width int) {
	for _, c := range formatComment(comment, width) {
		AddComment(commentList, c)
	}
}

// AddUnwrappedComments adds comments to the specified list. The text is not wrapped.
// Respects any existing line breaks specified by \n or <br>
func AddUnwrappedComments(commentList *dst.Decorations, comments []string) {
	for _, comment := range comments {
		// Skip empty comments
		if comment == "" {
			continue
		}

		AddUnwrappedComment(commentList, comment)
	}
}

// AddUnwrappedComment adds a single comment to the specified list. The text is not wrapped.
// Respects any existing line breaks specified by \n or <br>.
func AddUnwrappedComment(commentList *dst.Decorations, comment string) {
	for _, c := range formatComment(comment, 1000) { // We don't want line-wrapping with these comments
		AddComment(commentList, c)
	}
}

// AddComments adds preformatted comments to the specified list
func AddComments(commentList *dst.Decorations, comments []string) {
	for _, comment := range comments {
		// Skip empty comments
		if comment == "" {
			continue
		}

		AddComment(commentList, comment)
	}
}

// AddComment adds a single comment line to the specified list
func AddComment(commentList *dst.Decorations, comment string) {
	line := comment

	if !strings.HasPrefix(line, "//") {
		line = "// " + line
	}

	commentList.Append(line)
}

// formatComment splits the supplied comment string up ready for use as a documentation comment
func formatComment(comment string, width int) []string {
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

// docCommentWrap applies word wrapping to the specified width to the slice of strings, returning
// a new slice
func docCommentWrap(lines []string, width int) []string {
	var result []string
	for _, l := range lines {
		result = append(result, Wordwrap(l, width)...)
	}

	return result
}

// Wordwrap applies word wrapping to the specified string, returning a slice containing the lines.
func Wordwrap(text string, width int) []string {
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

// CommentLength returns the text length of the comments, including EoLN characters
func CommentLength(comments dst.Decorations) int {
	length := 0
	for _, l := range comments {
		length += len(l) + 1 // length including EoLN
	}

	return length
}
