/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package reporting

import (
	"fmt"
	"io"
)

const (
	blockIndent     = "│   "
	itemPrefix      = "├── "
	lastItemPrefix  = "└── "
	lastBlockIndent = "    "
)

// StructureReport represents a hierarchical dump of structural information
type StructureReport struct {
	line   string
	nested []*StructureReport
}

// NewStructureReport creates a new StructureReport
func NewStructureReport(line string) *StructureReport {
	return &StructureReport{line: line}
}

// Addf formats a new line in the report, returning a nested report for any additional information
func (sr *StructureReport) Addf(format string, a ...any) *StructureReport {
	result := &StructureReport{line: fmt.Sprintf(format, a...)}
	sr.nested = append(sr.nested, result)
	return result
}

func (sr *StructureReport) SaveTo(writer io.Writer) error {
	var indents []string

	return sr.writeBlock(writer, indents, "", "")
}

// writeTo writes a block of lines from this StructureReport to a writer
func (sr *StructureReport) writeBlock(
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

	_, err = io.WriteString(writer, sr.line)
	if err != nil {
		return err
	}

	_, err = io.WriteString(writer, "\n")
	if err != nil {
		return err
	}

	nested := append(indents, prefixForSubItems)
	for index, line := range sr.nested {
		ind := itemPrefix
		sub := blockIndent
		if index == len(sr.nested)-1 {
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
