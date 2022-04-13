/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package reporting

import (
	"fmt"
	"strings"
)

type MarkdownTable struct {
	content [][]string
	widths  []int
}

// NewMarkdownTable returns a new Markdown table with the specified columns
func NewMarkdownTable(columns ...string) *MarkdownTable {
	result := &MarkdownTable{}
	result.AddRow(columns...)
	return result
}

// AddRow adds an entire row to the table, tracking widths for final formatting
func (t *MarkdownTable) AddRow(row ...string) {
	t.content = append(t.content, row)
	for i, r := range row {
		w := len(r)
		if i >= len(t.widths) {
			t.widths = append(t.widths, w)
		} else if w > t.widths[i] {
			t.widths[i] = w
		}
	}
}

// WriteTo renders the Markdown table into the specified buffer
func (t *MarkdownTable) WriteTo(buffer *strings.Builder) {
	for i, r := range t.content {
		t.renderRow(r, buffer)
		if i == 0 {
			t.renderRowDivider(buffer)
		}
	}
}

// renderRow writes a single row into the buffer
func (t *MarkdownTable) renderRow(row []string, buffer *strings.Builder) {
	buffer.WriteRune('|')
	for i, c := range row {
		buffer.WriteString(fmt.Sprintf(" %*s |", -t.widths[i], c))
	}

	buffer.WriteString("\n")
}

// renderRowDivider writes a dividing line into the buffer
func (t *MarkdownTable) renderRowDivider(buffer *strings.Builder) {
	buffer.WriteString("|")
	for _, w := range t.widths {
		for i := -2; i < w; i++ {
			buffer.WriteRune('-')
		}

		buffer.WriteRune('|')
	}

	buffer.WriteString("\n")
}
