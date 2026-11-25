/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package reporting

import "io"

// MarkdownReport is a helper for producing reports in Markdown format
type MarkdownReport struct {
	content []string
}

// NewMarkdownReport creates a new instance of MarkdownReport
func NewMarkdownReport() *MarkdownReport {
	return &MarkdownReport{}
}

// WriteTo emits the report to the supplied writer
func (r *MarkdownReport) WriteTo(writer io.Writer) error {
	for _, line := range r.content {
		if _, err := io.WriteString(writer, line); err != nil {
			return err
		}
		if _, err := io.WriteString(writer, "\n"); err != nil {
			return err
		}
	}

	return nil
}

// AddLines adds one or more lines to the report
func (r *MarkdownReport) AddLines(lines ...string) {
	r.content = append(r.content, lines...)
}

// AddHeading adds a level 2 header to the report
func (r *MarkdownReport) AddHeading(header string) {
	r.AddLines(
		"",
		"## "+header,
		"",
	)
}

// AddSubheading adds a level 3 header to the report
func (r *MarkdownReport) AddSubheading(subheading string) {
	r.AddLines(
		"",
		"### "+subheading,
		"",
	)
}

// AddTable writes a MarkdownTable that's been prepared elsewhere
func (r *MarkdownReport) AddTable(table *MarkdownTable) {
	r.AddLines(table.ToStrings()...)
}
