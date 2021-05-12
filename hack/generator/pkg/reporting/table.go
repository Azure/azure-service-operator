/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package reporting

import (
	"fmt"
	"sort"
	"strings"
)

type Table struct {
	// Captions for each row
	rows []string
	// Length of longest row caption
	rowWidth int
	// Captions for each column
	cols []string
	// Width required for each column
	colWidths []int
	// Cell content, arranged per row, then per column
	cells map[string]map[string]string
}

func NewTable() *Table {
	return &Table{
		cells: make(map[string]map[string]string),
	}
}

// Rows returns a slice containing the captions of all the rows of the table
// A new slice is returned to avoid violations of encapsulation
func (t *Table) Rows() []string {
	var result []string
	result = append(result, t.rows...)
	return result
}

// AddRow adds the specified row to the table if it doesn't already exist
func (t *Table) AddRow(row string) {
	if t.indexOfRow(row) == -1 {
		t.rows = append(t.rows, row)
		if len(row) > t.rowWidth {
			t.rowWidth = len(row)
		}
	}
}

// SortRows allows rows to be sorted by caption
func (t *Table) SortRows(less func(top string, bottom string) bool) {
	sort.Slice(t.rows, func(i, j int) bool {
		return less(t.rows[i], t.rows[j])
	})
}

// Columns returns a slice containing the captions of all the columns of the table
// A new slice is returned to avoid violations of encapsulation
func (t *Table) Columns() []string {
	var result []string
	result = append(result, t.cols...)
	return result
}

// AddColumn adds the specified column to the table if it doesn't already exist
func (t *Table) AddColumn(col string) {
	index := t.indexOfColumn(col)
	if index == -1 {
		t.cols = append(t.cols, col)
		t.colWidths = append(t.colWidths, len(col))
	}
}

// SortColumns allows columns to be sorted by caption
func (t *Table) SortColumns(less func(left string, right string) bool) {
	sort.Slice(t.cols, func(i, j int) bool {
		return less(t.cols[i], t.cols[j])
	})
}

// SetCell sets the content of a given cell of the table
func (t *Table) SetCell(row string, col string, cell string) {
	t.AddColumn(col)
	t.AddRow(row)
	rowCells := t.getRowCells(row)
	rowCells[col] = cell

	index := t.indexOfColumn(col)
	if len(cell) > t.colWidths[index] {
		t.colWidths[index] = len(cell)
	}
}

func (t *Table) WriteTo(buffer *strings.Builder) {
	buffer.WriteString(t.renderHeader())
	buffer.WriteString(t.renderDivider())
	for _, r := range t.rows {
		buffer.WriteString(t.renderRow(r))
	}
}

func (t *Table) getRowCells(row string) map[string]string {
	if m, ok := t.cells[row]; ok {
		return m
	}

	result := make(map[string]string)
	t.cells[row] = result
	return result
}

func (t *Table) renderHeader() string {
	var result strings.Builder

	result.WriteString(fmt.Sprintf("| %*s |", t.rowWidth, ""))

	for _, c := range t.cols {
		result.WriteString(" ")
		result.WriteString(c)
		result.WriteString(" |")
	}

	result.WriteString("\n")
	return result.String()
}

func (t *Table) renderDivider() string {
	var result strings.Builder

	result.WriteString("|")
	for i := -2; i < t.rowWidth; i++ {
		result.WriteRune('-')
	}
	result.WriteRune('|')

	for w := range t.cols {
		width := t.colWidths[w]
		for i := -2; i < width; i++ {
			result.WriteRune('-')
		}

		result.WriteRune('|')
	}

	result.WriteString("\n")
	return result.String()
}

func (t *Table) renderRow(row string) string {
	var result strings.Builder
	cells := t.getRowCells(row)

	result.WriteString(fmt.Sprintf("| %*s |", -t.rowWidth, row))
	for i, c := range t.cols {
		content := cells[c]
		result.WriteString(fmt.Sprintf(" %*s |", -t.colWidths[i], content))
	}

	result.WriteString("\n")
	return result.String()
}

func (t *Table) indexOfRow(row string) int {
	for i, r := range t.rows {
		if r == row {
			return i
		}
	}

	return -1
}

func (t *Table) indexOfColumn(col string) int {
	for i, c := range t.cols {
		if c == col {
			return i
		}
	}

	return -1
}
