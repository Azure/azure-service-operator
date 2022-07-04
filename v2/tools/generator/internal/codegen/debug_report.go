/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package codegen

import (
	"io"
)

const (
	blockIndent     = "│   "
	itemPrefix      = "├── "
	lastItemPrefix  = "└── "
	lastBlockIndent = "    "
)

// debugReport represents a heirarchical dump of debug information
type debugReport struct {
	line   string
	nested []*debugReport
}

// newDebugReport creates a new debugReport
func newDebugReport(line string) *debugReport {
	return &debugReport{line: line}
}

// add adds a new debugReport nested in the current debugReport, returning the nested report
func (dr *debugReport) add(line string) *debugReport {
	result := &debugReport{line: line}
	dr.nested = append(dr.nested, result)
	return result
}

func (dr *debugReport) saveTo(writer io.Writer) error {
	var indents []string

	return dr.writeBlock(writer, indents, "", "")
}

// writeTo writes a block of lines from this debugReport to a writer
func (dr *debugReport) writeBlock(
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
