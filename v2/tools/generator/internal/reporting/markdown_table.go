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
	content    [][]string
	widths     []int
	alignments []ColumnAlignment
}

// ColumnAlignment specifies how a column of a MarkdownTable should be aligned
type ColumnAlignment int

const (
	// AlignDefault leaves alignment unspecified, allowing the renderer to choose
	AlignDefault ColumnAlignment = iota
	// AlignLeft explicitly left-aligns the column
	AlignLeft
	// AlignCenter centers the column
	AlignCenter
	// AlignRight right-aligns the column
	AlignRight
)

// NewMarkdownTable returns a new Markdown table with the specified columns
func NewMarkdownTable(columns ...string) *MarkdownTable {
	result := &MarkdownTable{}
	result.AddRow(columns...)
	return result
}

// SetAlignment specifies how the given (0-based) column should be aligned when rendered.
func (t *MarkdownTable) SetAlignment(column int, alignment ColumnAlignment) {
	for column >= len(t.alignments) {
		t.alignments = append(t.alignments, AlignDefault)
	}

	t.alignments[column] = alignment
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
		fmt.Fprintf(buffer, " %*s |", -t.widths[i], c)
	}

	buffer.WriteString("\n")
}

// renderRowDivider writes a dividing line into the buffer
func (t *MarkdownTable) renderRowDivider(buffer *strings.Builder) {
	buffer.WriteString("|")
	for i, w := range t.widths {
		left, right := false, false
		if i < len(t.alignments) {
			switch t.alignments[i] {
			case AlignLeft:
				left = true
			case AlignCenter:
				left, right = true, true
			case AlignRight:
				right = true
			case AlignDefault:
				// no colons
			}
		}

		// Total length matches the unaligned case (width+2); colons replace dashes as needed.
		dashes := w + 2
		if left {
			dashes--
		}

		if right {
			dashes--
		}

		if left {
			buffer.WriteRune(':')
		}

		for d := 0; d < dashes; d++ {
			buffer.WriteRune('-')
		}

		if right {
			buffer.WriteRune(':')
		}

		buffer.WriteRune('|')
	}

	buffer.WriteString("\n")
}
