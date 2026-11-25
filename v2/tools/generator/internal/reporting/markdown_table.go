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

// ToStrings renders the Markdown table as a slice of lines
func (t *MarkdownTable) ToStrings() []string {
	result := make([]string, 0, len(t.content)+1)
	for i, r := range t.content {
		result = append(result, t.renderRow(r))
		if i == 0 {
			result = append(result, t.renderRowDivider())
		}
	}

	return result
}

// WriteTo renders the Markdown table into the specified buffer
func (t *MarkdownTable) WriteTo(buffer *strings.Builder) {
	for _, line := range t.ToStrings() {
		buffer.WriteString(line)
		buffer.WriteRune('\n')
	}
}

// renderRow renders a single row
func (t *MarkdownTable) renderRow(row []string) string {
	var buffer strings.Builder
	buffer.WriteRune('|')
	for i, c := range row {
		alignment := AlignDefault
		if i < len(t.alignments) {
			alignment = t.alignments[i]
		}

		switch alignment {
		case AlignCenter:
			padding := t.widths[i] - len(c)
			left := padding / 2
			right := padding - left
			fmt.Fprintf(&buffer, " %*s%s%*s |", left, "", c, right, "")
		case AlignRight:
			fmt.Fprintf(&buffer, " %*s |", t.widths[i], c)
		default:
			fmt.Fprintf(&buffer, " %*s |", -t.widths[i], c)
		}
	}

	return buffer.String()
}

// renderRowDivider renders a dividing line
func (t *MarkdownTable) renderRowDivider() string {
	var buffer strings.Builder
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

	return buffer.String()
}
