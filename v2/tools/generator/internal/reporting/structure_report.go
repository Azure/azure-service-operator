/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package reporting

import (
	"io"
)

const (
	blockIndent     = "│   "
	itemPrefix      = "├── "
	lastItemPrefix  = "└── "
	lastBlockIndent = "    "
)

// StructureReport represents a heirarchical dump of structural information
type StructureReport struct {
	line   string
	nested []*StructureReport
}

// NewStructureReport creates a new StructureReport
func NewStructureReport(line string) *StructureReport {
	return &StructureReport{line: line}
}

// Add adds a new StructureReport nested in the current StructureReport, returning the nested report
func (dr *StructureReport) Add(line string) *StructureReport {
	result := &StructureReport{line: line}
	dr.nested = append(dr.nested, result)
	return result
}

func (dr *StructureReport) SaveTo(writer io.Writer) error {
	var indents []string

	return dr.writeBlock(writer, indents, "", "")
}

// writeTo writes a block of lines from this StructureReport to a writer
func (dr *StructureReport) writeBlock(
	writer io.Writer,
	indents []string,
	prefixForItem string,
	prefixForSubItems string) error {

	// Write existing prefix
	for _, i := range indents {
		_, err := io.WriteString(writer, i)
		if err != nil {
			return err
		}
	}

	_, err := io.WriteString(writer, prefixForItem)
	if err != nil {
		return err
	}

	_, err = io.WriteString(writer, dr.line)
	if err != nil {
		return err
	}

	_, err = io.WriteString(writer, "\n")
	if err != nil {
		return err
	}

	nested := append(indents, prefixForSubItems)
	for index, line := range dr.nested {
		ind := itemPrefix
		sub := blockIndent
		if index == len(dr.nested)-1 {
			ind = lastItemPrefix
			sub = lastBlockIndent
		}

		err = line.writeBlock(writer, nested, ind, sub)
		if err != nil {
			return err
		}
	}

	return nil
}
