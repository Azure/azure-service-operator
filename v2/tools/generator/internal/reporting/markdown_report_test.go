/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package reporting

import (
	"bytes"
	"testing"
)

func TestMarkdownReport_AddTable_AddsTableLines(t *testing.T) {
	t.Parallel()

	report := NewMarkdownReport()
	report.AddLines("Before")

	table := NewMarkdownTable("Name")
	table.AddRow("Alpha")
	report.AddTable(table)

	report.AddLines("After")

	var buffer bytes.Buffer
	if err := report.WriteTo(&buffer); err != nil {
		t.Fatal(err)
	}

	expected := "Before\n| Name  |\n|-------|\n| Alpha |\nAfter\n"
	if expected != buffer.String() {
		t.Fatalf("expected %q, got %q", expected, buffer.String())
	}
}
