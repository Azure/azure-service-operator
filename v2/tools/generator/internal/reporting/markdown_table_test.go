/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package reporting

import (
	"strings"
	"testing"

	"github.com/onsi/gomega"
	"github.com/sebdah/goldie/v2"
)

func TestMarkdownTable_ToStrings_ReturnsUnterminatedLines(t *testing.T) {
	t.Parallel()
	g := gomega.NewWithT(t)

	table := NewMarkdownTable("Name", "Status")
	table.SetAlignment(0, AlignLeft)
	table.SetAlignment(1, AlignCenter)
	table.AddRow("Alpha", "New")

	expected := []string{
		"| Name  | Status |",
		"|:------|:------:|",
		"| Alpha |  New   |",
	}

	g.Expect(table.ToStrings()).To(gomega.Equal(expected))
}

func TestMarkdownTable_GivesExpectedResults(t *testing.T) {
	t.Parallel()
	g := goldie.New(t)

	table := NewMarkdownTable("Name", "Description")
	table.AddRow("Alpha", "The first, be cautious")
	table.AddRow("Beta", "The second, should be stable")
	table.AddRow("Gamma", "The third, general availability")

	var buff strings.Builder
	table.WriteTo(&buff)

	g.Assert(t, t.Name(), []byte(buff.String()))
}

func TestMarkdownTable_WithAlignment_GivesExpectedResults(t *testing.T) {
	t.Parallel()
	g := goldie.New(t)

	table := NewMarkdownTable("Name", "Description", "Status")
	table.SetAlignment(0, AlignLeft)
	table.SetAlignment(1, AlignRight)
	table.SetAlignment(2, AlignCenter)
	table.AddRow("Alpha", "The first, be cautious", "New")
	table.AddRow("Beta", "The second, should be stable", "Modified")
	table.AddRow("Gamma", "The third, general availability", "-")

	var buff strings.Builder
	table.WriteTo(&buff)

	g.Assert(t, t.Name(), []byte(buff.String()))
}
