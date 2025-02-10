/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package reporting

import (
	"fmt"
	"io"
	"strings"
)

const (
	blockIndent     = "│   "
	itemPrefix      = "├── "
	lastItemPrefix  = "└── "
	lastBlockIndent = "    "
)

// StructureReport represents a hierarchical dump of structural information
type StructureReport struct {
	line   string             // a line of content to output
	isRoot bool               // true if this is the root line of a report
	nested []*StructureReport // nested content
}

// NewStructureReport creates a new StructureReport
func NewStructureReport(line string) *StructureReport {
	return &StructureReport{
		line:   line,
		isRoot: true,
	}
}

// Addf formats a new line in the report, returning a nested report for any additional information
func (sr *StructureReport) Addf(format string, a ...any) *StructureReport {
	result := &StructureReport{
		line: fmt.Sprintf(format, a...),
	}
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
	prefixForSubItems string,
) error {
	// Write our line
	err := sr.writeLine(writer, indents, prefixForItem+sr.line)
	if err != nil {
		return err
	}

	if sr.isRoot {
		// Write an underline under the heading
		underline := strings.Repeat("-", len(sr.line))
		err = sr.writeLine(writer, indents, prefixForItem+underline)
		if err != nil {
			return err
		}
	}

	indents = append(indents, prefixForSubItems)
	for index, line := range sr.nested {
		var ind string
		var sub string
		if sr.isRoot {
			// Items immediately under the root have no prefix
			ind = ""
			sub = ""
		} else if index == len(sr.nested)-1 {
			// Last items in the list are handled differently
			ind = lastItemPrefix
			sub = lastBlockIndent
		} else {
			// Regular indent
			ind = itemPrefix
			sub = blockIndent
		}

		err := line.writeBlock(writer, indents, ind, sub)
		if err != nil {
			return err
		}

		// Write a blank line between items if we're at the root
		if sr.isRoot {
			_, err = io.WriteString(writer, "\n")
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (sr *StructureReport) writeLine(
	writer io.Writer,
	indents []string,
	line string,
) error {
	// Write existing prefix
	for _, i := range indents {
		_, err := io.WriteString(writer, i)
		if err != nil {
			return err
		}
	}

	_, err := io.WriteString(writer, line)
	if err != nil {
		return err
	}

	_, err = io.WriteString(writer, "\n")
	if err != nil {
		return err
	}

	return nil
}
