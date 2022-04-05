/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package reporting

import (
	"sort"
	"strings"
)

type SparseTable struct {
	// title for the entire table
	title string
	// rows contains the captions for each row
	rows []string
	// rowWidth is the length of the longest row caption
	rowWidth int
	// cols contains the captions for each column
	cols []string
	// colWidths contains the width required for each column, indexed by caption
	colWidths map[string]int
	// cells is the content for each cell, arranged per row, then per column
	cells map[string]map[string]string
}

func NewSparseTable(title string) *SparseTable {
	return &SparseTable{
		title:     title,
		cells:     make(map[string]map[string]string),
		rowWidth:  len(title),
		colWidths: make(map[string]int),
	}
}

// Rows returns a slice containing the captions of all the rows of the table
// A new slice is returned to avoid violations of encapsulation
func (table *SparseTable) Rows() []string {
	var result []string
	result = append(result, table.rows...)
	return result
}

// AddRow adds the specified row to the table if it doesn't already exist
func (table *SparseTable) AddRow(row string) {
	if table.indexOfRow(row) == -1 {
		table.rows = append(table.rows, row)
		if len(row) > table.rowWidth {
			table.rowWidth = len(row)
		}
	}
}

// SortRows allows rows to be sorted by caption
func (table *SparseTable) SortRows(less func(top string, bottom string) bool) {
	sort.Slice(table.rows, func(i, j int) bool {
		return less(table.rows[i], table.rows[j])
	})
}

// Columns returns a slice containing the captions of all the columns of the table
// A new slice is returned to avoid violations of encapsulation
func (table *SparseTable) Columns() []string {
	var result []string
	result = append(result, table.cols...)
	return result
}

// AddColumn adds the specified column to the table if it doesn't already exist
func (table *SparseTable) AddColumn(col string) {
	index := table.indexOfColumn(col)
	if index == -1 {
		table.cols = append(table.cols, col)
		table.colWidths[col] = len(col)
	}
}

// SortColumns allows columns to be sorted by caption
func (table *SparseTable) SortColumns(less func(left string, right string) bool) {
	sort.Slice(table.cols, func(i, j int) bool {
		return less(table.cols[i], table.cols[j])
	})
}

// SetCell sets the content of a given cell of the table
func (table *SparseTable) SetCell(row string, col string, cell string) {
	table.AddColumn(col)
	table.AddRow(row)
	rowCells := table.getRowCells(row)
	rowCells[col] = cell

	if len(cell) > table.colWidths[col] {
		table.colWidths[col] = len(cell)
	}
}

func (table *SparseTable) WriteTo(buffer *strings.Builder) {
	headings := []string {table.title}
	headings = append(headings, table.cols...)

	mt := NewMarkdownTable(headings...)
	for _, r := range table.rows {
		row := table.createRow(r)
		mt.AddRow(row...)
	}

	mt.WriteTo(buffer)
}

func (table *SparseTable) getRowCells(row string) map[string]string {
	if m, ok := table.cells[row]; ok {
		return m
	}

	result := make(map[string]string)
	table.cells[row] = result
	return result
}

func (table *SparseTable) createRow(row string) []string {
	result := []string{row}
	cells := table.getRowCells(row)
	for _, c := range table.cols {
		content := cells[c]
		result = append(result, content)
	}

	return result
}

func (table *SparseTable) indexOfRow(row string) int {
	for i, r := range table.rows {
		if r == row {
			return i
		}
	}

	return -1
}

func (table *SparseTable) indexOfColumn(col string) int {
	for i, c := range table.cols {
		if c == col {
			return i
		}
	}

	return -1
}
